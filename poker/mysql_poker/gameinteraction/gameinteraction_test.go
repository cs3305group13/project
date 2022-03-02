package gameinteraction

import (
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/testing/utils"
	"github.com/cs3305/group13_2022/project/utils/env"
)


var envs = env.GetEnvironmentVariables("../../../testing.env")
	
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTableName = envs["TESTING_POKER_TABLES_TABLE"]


func TestPlayerRaised(t *testing.T) {
	
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	username := "derek"
	seatNumber := "1"
	raiseAmount := "10"
	
	playerRaised(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber, raiseAmount)
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

func TestTryTakeMoneyFromPlayer(t *testing.T) {

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	username := "derek"
	bid := "10.0"


	taken := TryTakeMoneyFromPlayer(DB, tx, testingPlayersTableName, tableID, username, bid)
	if taken == false {
		t.Error("Bid should have been accepted")
	}

	bid = "30.1"

	taken = TryTakeMoneyFromPlayer(DB, tx, testingPlayersTableName, tableID, username, bid)
	if taken == true {
		t.Error("Bid should not have been accepted")
	}
}

func TestPlayerFolded(t *testing.T) {
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	username := "barry"
	seatNumber := "4"
	
	PlayerFolded(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, username, seatNumber)
}

func TestGameInProgress(t *testing.T) {
	
	tableID := "1"
	w := utils.CreateRegularResponse()
	
	inProgress := GameInProgress(w, DB, testingTablesTableName, tableID)

	if ! inProgress {
		t.Error("Game should be in progress")
	}
}