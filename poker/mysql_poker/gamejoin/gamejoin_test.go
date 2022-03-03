package gamejoin

import (
	"testing"

	"github.com/cs3305/group13_2022/project/utils/env"
	"github.com/cs3305/group13_2022/project/mysql_db"
)

var envs = env.GetEnvironmentVariables("../../../testing.env")
	
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTableName = envs["TESTING_POKER_TABLES_TABLE"]



func TestCheckTableExists(t *testing.T) {
	
	tableID := "1"

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	exists := CheckTableExists(tx, testingTablesTableName, tableID)
	if exists == false {
		t.Error("TableID should match table row.")
	}
}



func TestUpdatePlayersSelectedGame(t *testing.T) {

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()
	
	tableID := "2"
	seatNumber := "2" 
	
	funds := UpdatePlayersSelectedGame(DB, tx, testingPlayersTableName, tableID, "derek", seatNumber)

	if funds != "30.00" {
		t.Errorf("Funds should be 30.0 and not %s", funds)
	}
}


