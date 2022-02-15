package Cards

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Card has the card suits and types in the deck
type Card struct {
	Type string
	Suit string
}

// Deck contains the cards to be shuffled and played
type Deck []Card

func CreateDeck() (deck Deck) {
	types := []string{"2", "3", "4", "5", "6", "7",
		"8", "9", "10", "Jack", "Queen", "King", "Ace"}

	suits := []string{"Club", "Spade", "Diamond", "Hearts"}

	// create the deck by looping and appending each suit and type to a card
	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type: types[i],
				Suit: suits[n],
			}
			deck = append(deck, card)
		}
	}
	return deck
}

// Function to shuffle the deck
func Shuffle(d Deck) Deck {
	for i := 0; i < len(d); i++ {

		randint, _ := rand.Int(rand.Reader, big.NewInt(52))

		// shuffles the deck by looping through 0 - length of deck, generates
		// a random number at the index of the loop and swap the random integer with the number at the index of the loop

		d[randint.Int64()], d[i] = d[i], d[randint.Int64()]

	}
	return d
}

// Function that prints the deck of cards to console
func Deal(d Deck, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(d[i])
	}
}
