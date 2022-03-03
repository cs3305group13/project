package poker

import (
	"html"
	"net/http"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils"
	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/utils/token"

	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gamecards"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameflow"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinteraction"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gamestart"
)



func HandleUserAjaxRequest( w http.ResponseWriter, r *http.Request ) {
	valid := token.TokenValid( w, r, true )
	if valid {
		envs := env.GetEnvironmentVariables("../../../production.env")
		handleUserButtons(w, r, envs)
	}
}

func handleUserButtons( w http.ResponseWriter, r *http.Request, envs map[string]string ) {

	r.ParseForm()
	if len(r.Form) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	DB := mysql_db.NewDB(envs)
	tablesTableName := envs["TABLES_TABLE"]
	playersTableName := envs["PLAYERS_TABLE"]
	pokerTablesTableName := envs["POKER_TABLES_TABLE"]

	readyButton := html.EscapeString(r.Form.Get("ready_button"))
	if readyButton != "" {
        gamestart.TryReadyUpPlayer(w, r, DB, tablesTableName, playersTableName, pokerTablesTableName)

		return
	} 
	tableID := token.GetTableID(r, "token")
	inProgress := gameinteraction.GameInProgress(w,DB, tablesTableName, tableID)
	if inProgress {

		username := token.GetUsername(r, "token")

		// checks if players turn.
		playersTurn := gameinteraction.PlayersTurn(DB, tablesTableName, playersTableName, tableID, username)
		if ! playersTurn {
			w.Write([]byte("MESSAGE:\nSorry not your turn"))
			return
		} else {

			db := mysql_db.EstablishConnection(DB)
			tx := mysql_db.NewTransaction(db)
			defer tx.Rollback()
			defer db.Close()

			seatNumber := token.GetSeatNumber(r, "token")

			action := html.EscapeString(r.Form.Get("action"))
			if action == "Fold" {
				gameinteraction.PlayerFolded(DB, tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber)
			} 

			amount := html.EscapeString(r.Form.Get("amount"))
			if ( action == "Raise" || action == "Call" || action == "Check" ) && amount != "" {
				gameinteraction.PlayerTakesAction( DB, tx, tablesTableName, playersTableName, pokerTablesTableName, tableID, username, seatNumber, amount )

				highestBidder, _ := gameinfo.GetHighestBidder(DB, pokerTablesTableName, tableID)
				nextAvailablePlayer := gameflow.NextAvailablePlayer(DB, playersTableName, tableID, username, seatNumber)
				if nextAvailablePlayer == highestBidder {
					gamecards.AddToCommunityCards(DB, tx, tablesTableName, pokerTablesTableName, tableID)
				}
			}


			err := tx.Commit()
			utils.CheckError(err)
		}
	}
}