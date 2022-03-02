package gameflow

import (
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils/env"
)


var envs = env.GetEnvironmentVariables("../../../testing.env")
	
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTableName = envs["TESTING_POKER_TABLES_TABLE"]

func TestNextAvailablePlayer(t *testing.T) {
	tableID := "1"
	
	NextAvailablePlayer(DB, testingPlayersTableName, tableID, "derek", "1")
}


func TestNextAvailablePlayers(t *testing.T) {
	tableID := "1"
	
	NextAvailablePlayers(DB, testingPlayersTableName, tableID, "derek", "1")
}

func TestSetNextAvailablePlayerAfterThisOne(t *testing.T) {
	
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	
	SetNextAvailablePlayerAfterThisOne(DB, tx, testingPokerTableName, testingPlayersTableName, tableID, "derek", "1", "dealer = ")
}