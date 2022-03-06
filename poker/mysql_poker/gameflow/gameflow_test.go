package gameflow

import (
	"fmt"
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/testing/mysql_poker"
	"github.com/cs3305/group13_2022/project/utils/env"
)

<<<<<<< HEAD
var envs = env.GetENvironemntVaribales("../../../testing.env")

=======

var envs = env.GetEnvironmentVariables("../../../testing.env")
	
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTableName = envs["TESTING_POKER_TABLES_TABLE"]


func TestClearUsersMoneyInPot(t *testing.T) {
<<<<<<< HEAD
	
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"

	ClearUsersMoneyInPot(DB, testingPlayersTableName, testingPokerTableName, tableID)
	
}

func TestNextAvailablePlayer(t *testing.T) {

	mysql_poker.RefreshPlayerTable( DB )

=======

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	
	ClearUsersMoneyInPot(DB, tx, testingPlayersTableName, testingPokerTableName, tableID)
}

func TestNextAvailablePlayer(t *testing.T) {

	mysql_poker.RefreshPlayerTable( DB )

>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
	tableID := "1"
	
	NextAvailablePlayer(DB, testingPlayersTableName, tableID, "derek", "1")
}


func TestNextAvailablePlayers(t *testing.T) {

	mysql_poker.RefreshPlayerTable( DB )

	tableID := "1"
	playerNames := NextAvailablePlayers(DB, testingPlayersTableName, tableID, "derek", "1")

	// No output because 5 seconds in where clause
	for i:=0; i<len(playerNames); i++ {
		fmt.Println(playerNames[i])
	}
}

func TestAssignThisPlayerToRole(t *testing.T) {

<<<<<<< HEAD
	tableID := "1"
	username := "derek"
	setOperation := fmt.Sprintf("current_player_making_move = '%s'", username)

	mysql_poker.RefreshTablesTable(DB)

	AssignThisPlayerToRole(DB, testingTablesTableName, tableID, username, setOperation)
=======
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	username := "derek"
	setOperation := fmt.Sprintf("current_player_making_move = '%s'", username)

	AssignThisPlayerToRole(tx, testingTablesTableName, tableID, username, setOperation)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
}

func TestSetNextAvailablePlayerAfterThisOne(t *testing.T) {
	mysql_poker.RefreshPlayerTable( DB )
<<<<<<< HEAD

	tableID := "1"
	
	SetNextAvailablePlayerAfterThisOne(DB, testingPokerTableName, testingPlayersTableName, tableID, "derek", "1", "dealer = ")
}

func TestUpdateUsersTimeSinceRequest(t *testing.T) {
=======

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	
	SetNextAvailablePlayerAfterThisOne(DB, tx, testingPokerTableName, testingPlayersTableName, tableID, "derek", "1", "dealer = ")
}

func TestUpdateUsersTimeSinceRequest(t *testing.T) {

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	username := "derek"
	tableID := "1"
	seatNumber := "1"
	
<<<<<<< HEAD
	UpdateUsersTimeSinceRequest(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, username, tableID, seatNumber)
=======
	UpdateUsersTimeSinceRequest(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, username, tableID, seatNumber)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
}