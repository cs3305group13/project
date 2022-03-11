package gamecards

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


func TestAddCards(t *testing.T) {
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"
	cardsToAdd := "2hTcAh"

	addCards(DB, testingPokerTableName, tableID, cardsToAdd)
}

func TestRefreshDeckAndCardsNotInDeck(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)

	deckString := "Ah2h3h4h5h6h7h8h9hThJhQhKhAd2d3d4d5d6d7d8d9dTdJdQdKdAs2s3s4s5s6s7s8s9sTsJsQsKsAc2c3c4c5c6c7c8c9cTcJcQcKc"
	cardsNotInDeckString := ""
	tableID := "1"

	refreshDeckAndCardsNotInDeck(DB, testingTablesTableName, deckString, cardsNotInDeckString, tableID)
}

func TestGetCards(t *testing.T) {

	tableID := "1"

	getCards(DB, testingTablesTableName, tableID)
}

func TestAssignPlayerHisCards(t *testing.T) {

	mysql_poker.RefreshPlayerTable(DB)

	tableID := "1"
	username := "derek"
	cardsString := "AHAC"  // ace of hearts and clubs

	assignPlayerHisCards(DB, testingPlayersTableName, tableID, username, cardsString)
}

func TestGivePlayersTheirCards(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)

	tableID := "1"
	
	GivePlayersTheirCards(DB, testingTablesTableName, testingPlayersTableName, tableID)
}


func TestAddToCommunityCards(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)
	
	tableID := "1"
	gameEndedEarly := false

	AddToCommunityCards(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, gameEndedEarly)

}

func TestAddToCommunityCards_gameEndedEarly(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"
	gameEndedEarly := true

	AddToCommunityCards(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, gameEndedEarly)
}

