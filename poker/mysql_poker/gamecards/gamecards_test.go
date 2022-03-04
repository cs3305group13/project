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


func TestAddCards( t *testing.T) {
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"
	cardsToAdd := "2h10CAH"

	addCards(DB, testingPokerTableName, tableID, cardsToAdd)
}


func TestRefreshDeckAndCardsNotInDeck(t *testing.T) {
	mysql_poker.RefreshTablesTable(DB)

	deckString := "AH2H3H4H5H6H7H8H9H10HJHQHKHAD2D3D4D5D6D7D8D9D10DJDQDKDAS2S3S4S5S6S7S8S9S10SJSQSKSAC2C3C4C5C6C7C8C9C10CJCKC"
	cardsNotInDeckString := "QC"
	tableID := "1"

	refreshDeckAndCardsNotInDeck(DB, testingTablesTableName, deckString, cardsNotInDeckString, tableID)
}

func TestGetCards(t *testing.T) {

	tableID := "1"

	getCards(DB, testingTablesTableName, tableID)
}
func TestAssignPlayerHisCard(t *testing.T) {
	mysql_poker.RefreshPlayerTable(DB)

	tableID := "1"
	username := "derek"
	cardString := "AHAC"

	assignPlayerHisCard(DB, testingPlayersTableName, tableID, username, cardsString)
}

func TestGivePlayerTheirCards(t *testing.T) {
	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefresjPlayerTables(DB)

	tableID := "1"

	GivePlayersTheirCards(DB, testingTablesTableName, testingPlayersTableName, tableID)
}

func TestAddToCommunityCards(t *testing.T) {

	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)


	tableID := "1"
	gameEndedEarly := false

	AddToCommunityCards(DB, tx, testingTablesTableName, testingPokerTableName, tableID, gameEndedEarly)

}

func TestAddToCommunityCards_gameEndedEarly(t *testing.T) {
	mysql_poker.RefreshTablesTable(DB)
	mysql_poker.RefreshPlayerTable(DB)
	mysql_poker.RefreshPokerTable(DB)

	tableID := "1"
	gameEndedEarly := true 

	AddToCommunityCards(DB, testingTablesTableName, testingPlayersTableName, testingPokerTableName, tableID, gameEndedEarly)
}

