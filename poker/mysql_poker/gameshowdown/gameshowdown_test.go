package gameshowdown

import (
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/testing/mysql_poker"
	"github.com/cs3305/group13_2022/project/utils/env"
)


var envs = env.GetEnvironmentVariables("../../../testing.env")
	
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTableName = envs["TESTING_POKER_TABLES_TABLE"]

func TestGameShowDown(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"

	ShowDown(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID)
}


func TestGetEndOfGameCommunityCards(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)
	
	tableID := "1"

	getEndOfGameCommunityCards(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID)
}