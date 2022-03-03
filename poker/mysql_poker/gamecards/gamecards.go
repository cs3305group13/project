package gamecards

import (
	"database/sql"
	"fmt"

	"github.com/cs3305/group13_2022/project/cards"
	"github.com/cs3305/group13_2022/project.mysql_db"
	"github.com/cs3305/group13_2022/project/utils"
)

func AddToCommunityCards(DB *my_sql.DB, tx *sql.Tx, tablesTableName, pokerTablesTableName, tableId string )
	deck, cardsNotInDeck := getCards(DB, tablesTableName, tableID)

	communityCards := getCommunityCards(DB, pokerTablesTableName, tableID)

	var cardsToAdd string
	if len(*communityCards) == 0 {
		for i:=0; i<3; i++ {
			cardsToAdd += cards.TakeCard(deck, cardsNotInDeck)
		}
	else if 5 > len(*communityCards) && len(*communityCards) >= 3 {
		cardsToAdd = cards.TakeCard(deck, cardsNotInDeck)
	} else {
		panic("Can only add cards if there are no cards and as long as there are no more than 4 cards. ")
	}

	// reassign deckString with deck without taken card
	deckString := cards.DeckString(deck)
	// reassign cardsNotInDeckString for the same purpose
	cardsNotInDeckString := cards.DeckString(cardsNotInDeck)

	refreshDeckAndCardsNotInDeck(tx, tablesTableName, deckString, cardsNotInDeckString, tableID )
	addCards(tx, pokerTablesTableName, tableID, cardsToAdd)

}

func getCommunityCards(DB *mysql_db.DB, pokerTablesTableName, tableID string) ( communityCards, *cards.Deck ) {
	db := mysql_db.EstablishConnection(DB)
	defer db.close()

	query := fmt.Sprintf(`SELECT community_cards
						  FROM %s
						  WHERE table_id = %s`, pokerTablesTableName, tableID)
	
	var communityCardsString string
	err := db.QeuryRow(query).Scan(&communityCardsString)
	utils.CheckError(err)

	communityCards = cards.ExtractDeck(communityCardsString)

	return communityCards
}


// Retrieves deck and cards_not_in_deck from tables table
func getCards(DB *mysql_db.DB, tablesTableName, tableID string) (deck, cardsNotInDeck *cards.Deck) {
	db := mysql_db.EstablishConnection(DB)
	defer db.close()

	query := fmt.Sprintf(`SELECT deck, cards_not_in_deck
						  FROM %s
						  WHERE table_id = %s;`, tablesTableName, tableID)
	
	var deckString string
	var cardsNotInDeckString string
	err := db.QeuryRow(query).Scan(&deckString, &cardsNotInDeckString)
	utils.CheckError(err)

	deck = cards.ExtractDeck(deckString)
	cardsNotInDeck = cards.ExtractDeck(cardsNotInDeckString)

	return deck, cardsNotInDeck
}

func refreshDeckAndCardsNotInDeck(tx *sql.Tx, tablesTableName, deckString, cardsNotInDeckString, tableID string) {
	// refresh state of deck and cards_not_in_deck in tables table
	query := fmt.Sprintf(`UPDATE %s
						 SET deck = "%s", 
						 	cards_not_in_deck = "%s
						 WHERE table_id = %s;`, tablesTableName, deckString, cardsNotInDeckString, tableID)
	res, err := tx.Exec(query)
	utils.CheckError(err)

	if utils.GetNumberOfRowsAffected(res) != 1 {
		panic("Exactly one row should have been affected")
	}
}

func addCards(tx *sql.Tx, pokerTablesTableName, tableID, cardsToAdd string ) {

	query := fmt.Sprintf(`UPDATE %s
						  SET community_cards = CONCAT(community_cards, "%s"
						  WHERE table_id = %s;`, pokerTablesTableName, cardsToAdd, tableID)
	
	res, err := tx.Exec(query)
	utils.CheckError(err)

	if utils.GetNumberOfRowsAffected(res) != 1 {
		panic("Exactly one row should have been affected")
	}
}

