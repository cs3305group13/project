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

func TestFindIdleImportantUser(t *testing.T) {

	tableID := "1"

	usernameFound := findIdleImportantUser(DB, testingPlayersTableName, tableID, "5", "derek")

	if usernameFound == "" {
		t.Error("This user should have been retrieved")
	}
}

func TestRemoveIdleUsers(t *testing.T) {

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"

	removeIdleUsers(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTablesTableName, tableID)
}

