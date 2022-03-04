package gameinteraction

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
	
	mysql_poker.RefreshPlayerTable(DB)

	username := "derek"

	playerAllIn(DB, testingPlayersTableName, username)
}

func TestTryTakeMoneyFromPlayer(t *testing.T) {

	tableID := "1"
	username := "derek"
	bid := "10.0"

	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	taken := TryTakeMoneyFromPlayer(DB, testingPlayersTableName, testingPokerTableName, tableID, username, bid)
	if taken == false {
		t.Error("Bid should have been accepted")
	}

	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	bid = "30.1"

	taken = TryTakeMoneyFromPlayer(DB, testingPlayersTableName, testingPokerTableName, tableID, username, bid)
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
	
	PlayerFolded(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, nextPlayerFoundBool)
}

func TestPlayerFoldedButIsLastPlayer(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"
	username := "barry"
	seatNumber := "4"

	nextPlayerFoundBool := false
	
	PlayerFolded(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, nextPlayerFoundBool)
}

// ++++++++++ PlayerTakesAction() testing ++++++++++
func TestPlayerTakesCallAction(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"
	username := "derek"
	seatNumber := "1"
	raiseAmount := "1.0"

	action := PlayerTakesAction(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)

	if action != "CALLED" {
		t.Error("According to table with id 1 the current highest bid should be 1.0")
	}

	// reset tabes for next test below
	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	raiseAmount = "2.00"

	action = PlayerTakesAction(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)

	if action == "CALLED" {
		t.Error("According to table with id 1 the current highest bid should be 1.0")
	}
}

func TestPlayerTakesRaiseAction(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"
	username := "derek"
	seatNumber := "1"
	raiseAmount := "2.0"

	action := PlayerTakesAction(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)

	if action != "RAISED" {
		t.Errorf("expectedOutput: 'RAISED' outputReceived: %s", action)
	}
}

func TestPlayerTakesAllInAction(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"
	username := "derek"
	seatNumber := "1"
	raiseAmount := "30.0"

	action := PlayerTakesAction(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)

	if action != "ALL_IN" {
		t.Errorf("expectedOutput: 'ALL_IN' outputReceived: %s", action)
	}
}

func TestPlayerTakesFoldedAction(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"
	username := "derek"
	seatNumber := "1"
	raiseAmount := "0.5"

	action := PlayerTakesAction(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)

	if action != "FOLDED" {
		t.Errorf("expectedOutput: 'FOLDED' outputReceived: %s", action)
	}
}