package gameobserver

import (
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils/env"
)

var envs = env.GetEnvironmentVariables("../../../testing.env")
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTablesTableName = envs["TESTING_POKER_TABLES_TABLE"]

func TestGameObserver( t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"

	GameObserver(DB, testingTablesTableName, testingPlayersTableName, testingPokerTablesTableName, tableID)
}

func TestFindIdleImportantUser(t *testing.T) {

	mysql_poker.RefreshPlayerTable(DB)

	tableID := "1"

	usernameFound := findIdleImportantUser(DB, testingPlayersTableName, tableID, "5", "derek")

	if usernameFound != "" {
		t.Error("This user should not have been retrieved")
	}
}

func TestRemoveIdleUsers(t *testing.T) {

	tableID := "1"

	removeIdleUsers(DB, testingTablesTableName, testingPlayersTableName, testingPokerTablesTableName, tableID)
}

