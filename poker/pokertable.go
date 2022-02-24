package poker

import (
	"html"
	"net/http"
	"strconv"
	"text/template"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gamecreate"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gamejoin"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameobserver"
	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/utils/token"
)

type details struct {
	Username string
	TableID string
	SeatNumber string
	Funds string
}

func HandlePokerTableRequest( w http.ResponseWriter, r *http.Request ) {
	if token.TokenValid( w, r, true ) {
		var d details
	    d.Username = token.GetUsername(r, "token")
		d.TableID = token.GetTableID(r, "token")
		d.SeatNumber = token.GetSeatNumber(r, "token")
		d.Funds = token.GetFunds(r, "token")
		
	    attachPokerPage(w, d)
	}
}

func HandlePokerForm( w http.ResponseWriter, r *http.Request ) (joinedGame bool) {

	r.ParseForm()
	if len(r.Form) == 0 {
		http.Redirect(w, r, "userpage", http.StatusMovedPermanently)
		return false
	}

	pokerOnlineTableCode := html.EscapeString(r.Form.Get("table_code"))  // table code input field

	tableCode, err := strconv.Atoi(pokerOnlineTableCode)
	if err != nil {
		tableCode = 0
	}

	envs := env.GetEnvironmentVariables("../../../production.env")
	if tableCode == 0 {
		joinedGame = createOnlinePoker(w, r, envs)
	}
	if tableCode > 0 {
		joinedGame = joinOnlinePoker(w, r, envs, pokerOnlineTableCode)
	}
	return joinedGame
}


func createOnlinePoker( w http.ResponseWriter, r *http.Request, envs map[string]string ) (joinedGame bool) {
	
	username := token.GetUsername(r, "token")

	DB := mysql_db.NewDB(envs)
	playersTableName := envs["PLAYERS_TABLE"]
	tablesTableName := envs["TABLES_TABLE"]
	pokerTablesTableName := envs["POKER_TABLES_TABLE"]

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := gamecreate.AssignNewTable( tx, tablesTableName, pokerTablesTableName, username )

	seatNumber, _ := gameinfo.GetNextAvailableSeat(DB, playersTableName, tableID)
	
	funds := gamejoin.UpdatePlayersSelectedGame(tx, playersTableName, tableID, username, seatNumber)


	claims, expirationTime := token.NewDefaultClaims(username, "poker", tableID, seatNumber, funds)
	token.AppendTokenCookie(w, "token", claims, expirationTime)
	
	tx.Commit()

	go gameobserver.GameObserver(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID)

	return true
	
}

func joinOnlinePoker( w http.ResponseWriter, r *http.Request, envs map[string]string, tableID string ) (joinedGame bool) {
	
	DB := mysql_db.NewDB(envs)
	tablesTableName := envs["TABLES_TABLE"]
	playersTableName := envs["PLAYERS_TABLE"]

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	if gamejoin.CheckTableExists(tx, tablesTableName, tableID) == false {
		attachUserPage(w, "Sorry please make sure to specify a valid tableCode.")
		return false
	} else {
		username := token.GetUsername(r, "token")

		numberOfPlayersAtTable := gameinfo.GetNumberOfPlayersAtTable(DB, tablesTableName, tableID)

		maxNumberOfPlayers, _ := strconv.Atoi(envs["MAX_NUMBER_OF_PLAYERS"])
		if numberOfPlayersAtTable >= maxNumberOfPlayers {
			attachUserPage(w, "Sorry this table is full.")
			return false
		} else {
			seatNumber, seatFound := gameinfo.GetNextAvailableSeat(DB, playersTableName, tableID)
			if seatFound == false {
				attachUserPage(w, "Sorry this table is full.")
			    return false
			} else {
				funds := gamejoin.UpdatePlayersSelectedGame( tx, playersTableName, tableID, username, seatNumber )

				claims, expirationTime := token.NewDefaultClaims(username, "poker", tableID, seatNumber, funds)
				token.AppendTokenCookie(w, "token", claims, expirationTime)
				
				tx.Commit()
				return true
			}
		}
	}
}


func attachPokerPage( w http.ResponseWriter, d details ) {
	var tpl *template.Template
	tpl, _ = template.ParseGlob("templates/private/pokertable.html")
	tpl.ExecuteTemplate(w, "pokertable.html", d)
}

func attachUserPage(w http.ResponseWriter, message string) {
	var tpl *template.Template
	tpl, _ = template.ParseGlob("templates/private/userpage.html")
	tpl.ExecuteTemplate(w, "userpage.html", message)
}