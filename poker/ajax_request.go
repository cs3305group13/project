package poker

import (
	"net/http"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gamecontent"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameflow"
	"github.com/cs3305/group13_2022/project/utils"
	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/utils/token"
)


func HandleContentAjaxRequest(w http.ResponseWriter, r *http.Request) {
	access := token.TokenValid( w, r, true )
	if access == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	envs := env.GetEnvironmentVariables("../../../production.env")

	DB := mysql_db.NewDB(envs)
	playersTableName := envs["PLAYERS_TABLE"]
	tablesTableName := envs["TABLES_TABLE"]
	pokerTablesTableName := envs["POKER_TABLES_TABLE"]

	tableID := token.GetTableID( r, "token" )
	username := token.GetUsername( r, "token" )
	seatNumber := token.GetSeatNumber( r, "token" )

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	gameflow.UpdateUsersTimeSinceRequest(DB, tx, tablesTableName, playersTableName, pokerTablesTableName, username, tableID, seatNumber)

	err := tx.Commit()
	utils.CheckError(err)
	
	// retrieves the information about players and about which player is currently making a move.
	gameDetails := gamecontent.JSONGameDetails(DB, tablesTableName, playersTableName, pokerTablesTableName, tableID, username )

	w.Write(gameDetails)
}