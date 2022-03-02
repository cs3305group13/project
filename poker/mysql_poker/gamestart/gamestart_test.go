package gamestart

import (
	"testing"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/testing/mysql_poker"
	"github.com/cs3305/group13_2022/project/testing/utils"
	"github.com/cs3305/group13_2022/project/utils/env"
)


var envs = env.GetEnvironmentVariables("../../../testing.env")
	
var w = utils.CreateRegularResponse()
var r = utils.CreateRequestWithPokerCookie()

var DB = mysql_db.NewDB(envs)

var testingTablesTableName = envs["TESTING_TABLES_TABLE"]
var testingPlayersTableName = envs["TESTING_PLAYERS_TABLE"]
var testingPokerTablesTableName = envs["TESTING_POKER_TABLES_TABLE"]


func TestFindWhoShouldBeSmallAndBigBlind(t *testing.T) {
	mysql_poker.RefreshPlayers(DB, testingPlayersTableName)

	tableID := "1"
	currentPlayerMakingMove := "derek"
	seatNumber := "1"

	small, big, newCurrent := findWhoShouldBeSmallAndBigBlind(DB, testingPlayersTableName, tableID, currentPlayerMakingMove, seatNumber)

	if small == "" {
		t.Error("small blind wasn't retrieved")
	}
	if big == "" {
		t.Error("big blind wasn't retrieved")
	}
	if newCurrent == "" {
		t.Error("new current player wasn't retrieved")
	}
}

func TestBeginGame(t *testing.T) {

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"

	mysql_poker.RefreshPlayers(DB, testingPlayersTableName)

	beginGame(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTablesTableName, tableID)
}

func TestReadyUpPlayer(t *testing.T) {
	
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	username := "derek"

	readyUpPlayer(w, DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTablesTableName, tableID, username)
}

func TestStartGame(t *testing.T) {
	
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"

	startGame(tx, testingTablesTableName, testingPlayersTableName, testingPokerTablesTableName, tableID)
}