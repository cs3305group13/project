package gameinteraction

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

func TestPlayerChecked(t *testing.T) {
	tableID := "1"
	username := "derek"

	playerChecked(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username)
}

func TestPlayerRaised(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"
	username := "derek"
	seatNumber := "1"
	raiseAmount := "10.0"
	
	playerRaised(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)
}

func TestPlayersTurn(t *testing.T) {
	
	username := "barry"
	tableID := "1"

	playersTurn := PlayersTurn(DB, testingTablesTableName, testingPlayersTableName, tableID, username)
	if playersTurn == false {
		t.Error("players turn should be true")
	}

	username = "derek"
	tableID = "1"

	playersTurn = PlayersTurn(DB, testingTablesTableName, testingPlayersTableName, tableID, username)
	if playersTurn == true {
		t.Error("players turn should be false")
	}
}

func TestPlayerAllIn(t *testing.T) {
<<<<<<< HEAD
	
	mysql_poker.RefreshPlayerTable(DB)
=======

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	username := "derek"


	playerAllIn(tx, testingPlayersTableName, username)
}

func TestTryTakeMoneyFromPlayer(t *testing.T) {
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	username := "derek"

	playerAllIn(DB, testingPlayersTableName, username)
}

func TestTryTakeMoneyFromPlayer(t *testing.T) {

	tableID := "1"
	username := "derek"
	bid := "10.0"

	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

<<<<<<< HEAD
	taken := TryTakeMoneyFromPlayer(DB, testingPlayersTableName, testingPokerTableName, tableID, username, bid)
=======
	taken := TryTakeMoneyFromPlayer(DB, tx, testingPlayersTableName, testingPokerTableName, tableID, username, bid)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
	if taken == false {
		t.Error("Bid should have been accepted")
	}

	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	bid = "30.1"

<<<<<<< HEAD
	taken = TryTakeMoneyFromPlayer(DB, testingPlayersTableName, testingPokerTableName, tableID, username, bid)
=======
	taken = TryTakeMoneyFromPlayer(DB, tx, testingPlayersTableName, testingPokerTableName, tableID, username, bid)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
	if taken == true {
		t.Error("Bid should not have been accepted")
	}

}

func TestPlayerFolded(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"
	username := "barry"
	seatNumber := "4"

	nextPlayerFoundBool := true
	
<<<<<<< HEAD
	PlayerFolded(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, nextPlayerFoundBool)
}

func TestPlayerFoldedButIsLastPlayer(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)
=======
	PlayerFolded(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, nextPlayerFoundBool)
}

func TestPlayerFoldedButIsLastPlayer(t *testing.T) {
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	tableID := "1"
	username := "barry"
	seatNumber := "4"

	nextPlayerFoundBool := false
	
<<<<<<< HEAD
	PlayerFolded(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, nextPlayerFoundBool)
=======
	PlayerFolded(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, nextPlayerFoundBool)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
}

// ++++++++++ PlayerTakesAction() testing ++++++++++
func TestPlayerTakesCallAction(t *testing.T) {

<<<<<<< HEAD
	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)
=======
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	tableID := "1"
	username := "derek"
	seatNumber := "1"
	raiseAmount := "1.0"

<<<<<<< HEAD
	action := PlayerTakesAction(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)
=======
	action := PlayerTakesAction(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	if action != "CALLED" {
		t.Error("According to table with id 1 the current highest bid should be 1.0")
	}

<<<<<<< HEAD
	// reset tabes for next test below
	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	raiseAmount = "2.00"

	action = PlayerTakesAction(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)

	if action == "CALLED" {
=======
	raiseAmount = "1.99"

	action = PlayerTakesAction(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)

	if action != "CALLED" {
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c
		t.Error("According to table with id 1 the current highest bid should be 1.0")
	}
}

func TestPlayerTakesRaiseAction(t *testing.T) {

<<<<<<< HEAD
	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)
=======
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	tableID := "1"
	username := "derek"
	seatNumber := "1"
	raiseAmount := "2.0"

<<<<<<< HEAD
	action := PlayerTakesAction(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)
=======
	action := PlayerTakesAction(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	if action != "RAISED" {
		t.Errorf("expectedOutput: 'RAISED' outputReceived: %s", action)
	}
}

func TestPlayerTakesAllInAction(t *testing.T) {

<<<<<<< HEAD
	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)
=======
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	tableID := "1"
	username := "derek"
	seatNumber := "1"
	raiseAmount := "30.0"

<<<<<<< HEAD
	action := PlayerTakesAction(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)
=======
	action := PlayerTakesAction(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	if action != "ALL_IN" {
		t.Errorf("expectedOutput: 'ALL_IN' outputReceived: %s", action)
	}
}

func TestPlayerTakesFoldedAction(t *testing.T) {

<<<<<<< HEAD
	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)
=======
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	tableID := "1"
	username := "derek"
	seatNumber := "1"
	raiseAmount := "0.5"

<<<<<<< HEAD
	action := PlayerTakesAction(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)
=======
	action := PlayerTakesAction(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)
>>>>>>> ecc4f5f74a4a414e36a17abc4e3f6d391559f80c

	if action != "FOLDED" {
		t.Errorf("expectedOutput: 'FOLDED' outputReceived: %s", action)
	}
}