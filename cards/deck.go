package cards

import (
	"crypto/rand"
	"math/big"
)

var availableValues = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var availableSuits = []string{"H", "D", "S", "C"}

type Deck []string

func NewDeck( numberOfDecks int ) (*Deck)  {
	var deck Deck
	for i := 0; i < numberOfDecks; i++ {
		for suit := 0; suit < 4; suit ++ {
			for value := 0; value < 13; value ++ {
				cardSuit := availableSuits[suit]
				cardValue := availableValues[value]

				deck = append(deck, cardValue + cardSuit)
			}
		}
	} 

	return &deck
}

func ExtractDeck( deckString string ) (*Deck) {
	var deck Deck
	var value string
	var suit string
	for i:=0; i<len(deckString); i++ {
		value = string(deckString[i])
		if value == "1" {   // if we iterate to the number 1 this means we are dealing with value 10, 10 is the only card value consisting of 2numbers 
			i ++
			value += string(deckString[i])
		}
		i ++
		suit = string(deckString[i])
		deck = append(deck, value+suit)
	}

	return &deck
}

// Function shuffles the supplied deck.
func Shuffle( deck *Deck ) {
	numOfRounds := 100
	for i := 0; i < numOfRounds; i ++ {
		for n := 0; n < len(*deck); n++ {
			randomIndex, _ := rand.Int(rand.Reader, big.NewInt( int64(len(*deck)) ))

			(*deck)[n], (*deck)[randomIndex.Int64()] = (*deck)[randomIndex.Int64()], (*deck)[n]
		}
	}
}

// Function removes card from top of deck, appends it into cardsNotInDeck and returns the card removed.
func TakeCard( deck, cardsNotInDeck *Deck ) (topCard string) {
	if len(*deck) <= 0 {
		return 
	}

	topCard = (*deck)[len(*deck)-1]    // card taken.

	*deck = (*deck)[:len(*deck) - 1]  // deck minus card taken.
	*cardsNotInDeck = append((*cardsNotInDeck), topCard)

	return topCard
}

// Method reinserts cards not in deck back into the original deck slice.
func RefillDeck( deck, cardsNotInDeck *Deck ) {
	*deck = append(*deck, *cardsNotInDeck...)
	*cardsNotInDeck = nil
}


func DeckString(deck *Deck) string {
	outputString := ""
	for i:=0; i<len(*deck); i++ {
		outputString += (*deck)[i]
	}
	return outputString
}