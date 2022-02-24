package gameinteraction

import (
	"database/sql"
	"net/http"

	"github.com/cs3305/group13_2022/project/mysql_db"
)

func GameInProgress(w http.ResponseWriter, r *http.Request, tx *sql.Tx, tablesTableName string) bool {
	
	return false
}


func TakeMoneyFromPlayer(DB *mysql_db.DB, playersTableName, pokerTablesTableName, tableID, smallBlind, smallBlindAmount string) {

}

func PlayersTurn(w http.ResponseWriter, r *http.Request, DB *mysql_db.DB, tablesTableName string) bool {

	return false
}


func PlayerFolded(w http.ResponseWriter, r *http.Request, DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName string) {

}


func PlayerRaised(w http.ResponseWriter, r *http.Request, DB *mysql_db.DB, tablesTableName, playersTableName, raiseAmount string)