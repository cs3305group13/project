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

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	cardsToAdd := "2HTCAH"

	addCards(tx, testingPokerTableName, tableID, cardsToAdd)
}

func TestRefreshDeckAndCardsNotInDeck(t *testing.T) {
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	deckString := "AH2H3H4H5H6H7H8H9HTHJHQHKHAD2D3D4D5D6D7D8D9DTDJDQDKDAS2S3S4S5S6S7S8S9STSJSQSKSAC2C3C4C5C6C7C8C9CTCJCQCKC"
	cardsNotInDeckString := ""
	tableID := "1"

	refreshDeckAndCardsNotInDeck(tx, testingTablesTableName, deckString, cardsNotInDeckString, tableID)
}

func TestGetCards(t *testing.T) {

	tableID := "1"

	getCards(DB, testingTablesTableName, tableID)
}

func TestAssignPlayerHisCards(t *testing.T) {
	
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	username := "derek"
	cardsString := "AHAC"  // ace of hearts and clubs

	assignPlayerHisCards(tx, testingPlayersTableName, tableID, username, cardsString)
}

func TestGivePlayersTheirCards(t *testing.T) {
	
	mysql_poker.RefreshPlayerTable( DB )

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	
	GivePlayersTheirCards(DB, tx, testingTablesTableName, testingPlayersTableName, tableID)
}


func TestAddToCommunityCards(t *testing.T) {

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()
	
	tableID := "1"
	gameEndedEarly := false

	AddToCommunityCards(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, gameEndedEarly)

}

func TestAddToCommunityCards_gameEndedEarly(t *testing.T) {

	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	tableID := "1"
	gameEndedEarly := true

	AddToCommunityCards(DB, tx, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, gameEndedEarly)
}

