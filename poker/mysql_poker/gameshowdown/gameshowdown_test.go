package gameshowdown

import (
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
<<<<<<< HEAD
	"github.com/cs3305/group13_2022/project/testing/mysql_poker"
=======
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
	"github.com/cs3305/group13_2022/project/utils/env"
)


var envs = env.GetEnvironmentVariables("../../../testing.env")
	
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTableName = envs["TESTING_POKER_TABLES_TABLE"]

func TestGameShowDown(t *testing.T) {
<<<<<<< HEAD

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"

	ShowDown(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID)
=======
	tableID := "1"

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	ShowDown(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
}


func TestGetEndOfGameCommunityCards(t *testing.T) {
<<<<<<< HEAD

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)
	
	tableID := "1"

	getEndOfGameCommunityCards(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID)
=======
	tableID := "1"


	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	getEndOfGameCommunityCards(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
}