package gamecards

import (
	"fmt"

	"github.com/cs3305/group13_2022/project/cards"
	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameflow"
	"github.com/cs3305/group13_2022/project/poker/mysql_poker/gameinfo"
	"github.com/cs3305/group13_2022/project/utils"
)


func AddToCommunityCards(DB *mysql_db.DB, tablesTableName, playersTableName, pokerTablesTableName, tableID string, gameEndedEarly bool) bool {
	deck, cardsNotInDeck := getCards(DB, tablesTableName, tableID)

	communityCards := gameinfo.GetCommunityCards(DB, pokerTablesTableName, tableID)

	var cardsToAdd string
	if len(*communityCards) == 0 {
		for i:=0; i<3; i++ {
			cardsToAdd += cards.TakeCard(deck, cardsNotInDeck)
		}

	} else if gameEndedEarly {
		i := len(*communityCards) - 1  // '-1' to offset for correct index
		if i < 0 {
			i = 0
		}
		
		for i < 5 {
			i += 1
			cardsToAdd += cards.TakeCard(deck, cardsNotInDeck)
		}
		return false

	} else if 5 > len(*communityCards) && len(*communityCards) >= 3  {
		cardsToAdd = cards.TakeCard(deck, cardsNotInDeck)

	} else {
		return false
	}

    // reassign deckString with deck without taken card
	deckString := cards.DeckString(deck)
	
	// reassign cardsNotInDeckString for the same purpose
	cardsNotInDeckString := cards.DeckString(cardsNotInDeck)

	refreshDeckAndCardsNotInDeck(DB, tablesTableName, deckString, cardsNotInDeckString, tableID)
	
	addCards(DB, pokerTablesTableName, tableID, cardsToAdd)
	
	return true
}


// function distributes a pair of cards to each player who is playing
func GivePlayersTheirCards(DB *mysql_db.DB, tablesTableName, playersTableName, tableID string) {
	deck, cardsNotInDeck := getCards(DB, tablesTableName, tableID)

	cardsNotInDeckAsString := cards.DeckString(cardsNotInDeck)
	if len(cardsNotInDeckAsString) > 0 {
		panic("Cards not in deck should be empty")
	}

	playerMakingMove, theirSeatNumber := gameinfo.GetCurrentPlayerMakingMove(DB, tablesTableName, playersTableName, tableID)
	playersPlaying := gameflow.NextAvailablePlayers(DB, playersTableName, tableID, playerMakingMove, theirSeatNumber)


	var card1 string
	var card2 string
	var playersCard string

	for i:=0; i<len(playersPlaying); i++ {
		card1 = cards.TakeCard(deck, cardsNotInDeck)
		card2 = cards.TakeCard(deck, cardsNotInDeck)

		playersCard = card1+card2

		playerName := playersPlaying[i]

		assignPlayerHisCards(DB, playersTableName, tableID, playerName, playersCard)
	}

	deckString := cards.DeckString(deck)
	cardsNotInDeckString := cards.DeckString(cardsNotInDeck)

	refreshDeckAndCardsNotInDeck(DB, tablesTableName, deckString, cardsNotInDeckString, tableID)
}

func assignPlayerHisCards(DB *mysql_db.DB, playersTableName, tableID, username, cardsString string) {
	
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()
	
	query := fmt.Sprintf(`UPDATE %s
	                      SET player_cards = "%s"
						  WHERE username = "%s";`, playersTableName, cardsString, username)


	_, err := db.Exec(query)

	utils.CheckError(err)
}

// Retrieves deck and cards_not_in_deck from tables table
func getCards(DB *mysql_db.DB, tablesTableName, tableID string) (deck, cardsNotInDeck *cards.Deck) {
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

	query := fmt.Sprintf(`SELECT deck, cards_not_in_deck
	                      FROM %s
						  WHERE table_id = %s;`, tablesTableName, tableID)

	var deckString string
	var cardsNotInDeckString string
	err := db.QueryRow(query).Scan(&deckString, &cardsNotInDeckString)
	utils.CheckError(err)

	deck = cards.ExtractDeck(deckString)
	cardsNotInDeck = cards.ExtractDeck(cardsNotInDeckString)

	return deck, cardsNotInDeck
}

func refreshDeckAndCardsNotInDeck(DB *mysql_db.DB, tablesTableName, deckString, cardsNotInDeckString, tableID string) {
	
	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

    // refresh state of deck and cards_not_in_deck in tables table
	query := fmt.Sprintf(`UPDATE %s
	                      SET deck = "%s",
						      cards_not_in_deck = "%s"
						  WHERE table_id = %s;`, tablesTableName, deckString, cardsNotInDeckString, tableID)
	res, err := db.Exec(query)
	utils.CheckError(err)

	if mysql_db.GetNumberOfRowsAffected(res) > 1 {
		panic("Exactly one row should have been affected")
	}

}


// appends cards to 'community cards' in 'poker_tables'
func addCards(DB *mysql_db.DB, pokerTablesTableName, tableID, cardsToAdd string) {

	db := mysql_db.EstablishConnection(DB)
	defer db.Close()

    query := fmt.Sprintf(`UPDATE %s
	                      SET community_cards = CONCAT(community_cards, "%s")
					      WHERE table_id = %s;`, pokerTablesTableName, cardsToAdd, tableID)
	
    res, err := db.Exec(query)
	utils.CheckError(err)

	if mysql_db.GetNumberOfRowsAffected(res) > 1 {
		panic("Exactly one row should have been affected")
	}
}