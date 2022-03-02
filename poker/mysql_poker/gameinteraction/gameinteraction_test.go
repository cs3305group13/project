package gameinteraction

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

func TestTryTakeMoneyFromPlayer(t *testing.T) {

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	username := "derek"
	bid := "10.0"


	taken := TryTakeMoneyFromPlayer(DB, tx, testingPlayersTableName, testingPokerTableName, tableID, username, bid)
	if taken == false {
		t.Error("Bid should have been accepted")
	}

	bid = "30.1"

	taken = TryTakeMoneyFromPlayer(DB, tx, testingPlayersTableName, testingPokerTableName, tableID, username, bid)
	if taken == true {
		t.Error("Bid should not have been accepted")
	}
}