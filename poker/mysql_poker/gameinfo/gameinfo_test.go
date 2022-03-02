package gameinfo

import (
	"fmt"
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/utils/env"
)



var envs = env.GetEnvironmentVariables("../../../testing.env")
	
var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTableName = envs["TESTING_POKER_TABLES_TABLE"]

func TestGetNumberOfPlayersAtTable(t *testing.T) {
	
	tableID := "1"

	numberOfPlayers := GetNumberOfPlayersAtTable( DB, testingPlayersTableName, tableID )

	if numberOfPlayers != 8 {
		t.Error("There should be exactly two players at table, first check there are two players bounded to table with id 1.")
	}
}

func TestGetNextAvailableSeat(t *testing.T) {
	
	tableID := "1"

	seatNumber, seatFound := GetNextAvailableSeat(DB, testingPlayersTableName, tableID)
	if seatFound {
		t.Error("There should be 8 users at this table already")
	}

	tableID = "2"
	seatNumber2, seatFound2 := GetNextAvailableSeat(DB, testingPlayersTableName, tableID)
	if ! seatFound2 {
		t.Error("There should be 1 users at this table already")
	}

	fmt.Println(seatNumber)
	fmt.Println(seatNumber2)
}

func TestGetPlayersFunds(t *testing.T) {
	username := "derek"

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()
	
	funds := GetPlayersFunds(tx, testingPlayersTableName, username)

	if funds != "30.00" {
		t.Errorf("%s should have 30.00 funds", username)
	}
	fmt.Println(funds)
}

func TestGetDealerAndHighestBidder(t *testing.T) {
	tableID := "1"
	
	highestBidder, highestBidderSeatNumber, dealer, dealerSeatNumber := GetDealerAndHighestBidder(DB, testingPlayersTableName, testingPokerTableName, tableID)

	if highestBidder == "" {
		t.Error("highestBidder wasn't retrieved correctly")
	}
	if highestBidderSeatNumber == "" {
		t.Error("highestBidderSeatNumber wasn't retrieved correctly")
	}
	if dealer == "" {
		t.Error("dealer wasn't retrieved correctly")
	}
	if dealerSeatNumber == "" {
		t.Error("dealerSeatNumber wasn't retrieved correctly")
	}
}

func TestGetCurrentPlayerMakingMove(t *testing.T) {
	tableID := "1"

	currentPlayerMakingMove, seatNumber := GetCurrentPlayerMakingMove(DB, testingTablesTableName, testingPlayersTableName, tableID)

	
	if currentPlayerMakingMove == "" {
		t.Error("current player making move name wasn't retrieved correctly")
	}

	if seatNumber == "" {
		t.Error("current player making move seat number wasn't retrieved correctly")
	}
}