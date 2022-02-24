package cards

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	sizeOfDeck := 1
	numberOfCardsPerDeck := 52

	deck := NewDeck(sizeOfDeck)

	if len(*deck) != sizeOfDeck * numberOfCardsPerDeck {
		t.Error("There is a logic error in NewDeck(sizeOfDeck)")
	}
}

func TestExtractDeck(t *testing.T) {
	deck := NewDeck(1)
	deckString := DeckString(deck)


	deckAfterBeingString := ExtractDeck(deckString)

	if ! reflect.DeepEqual(deck, deckAfterBeingString) {
		t.Error("tables are not deeply equal.")
	}

    for i:=0; i<len(*deck) && i<len(*deckAfterBeingString); i++ {
		if (*deck)[i] != (*deckAfterBeingString)[i] {
			t.Error("decks do not match")
		}
	}
}

func TestShuffle(t *testing.T) {
	deck := NewDeck(1)
	sortedDeck := NewDeck(1)

	Shuffle(deck)
	
	if reflect.DeepEqual(deck, sortedDeck) {
		t.Error("Slices should not be the same.")
	}
	fmt.Println(deck)
	fmt.Println(sortedDeck)
}

func TestTakeCard(t *testing.T) {
	deck := NewDeck(1)
	emptyDeck := NewDeck(0)

	card := TakeCard(deck, emptyDeck)

	// fmt.Println(card)
	// fmt.Println(deck)
	// fmt.Println(emptyDeck)

	if card == "" {
		t.Error("card wasn't given")
	}
	if len(*deck) != 51 {
		t.Error("card wasn't removed from deck")
	}
	if len(*emptyDeck) != 1 {
		t.Error("card wasn't added to emptyDeck.")
	}
}

func TestRefillDeck(t *testing.T) {
	deck := NewDeck(1)
	emptyDeck := NewDeck(0)

	TakeCard(deck, emptyDeck)

	RefillDeck(deck, emptyDeck)
	if len(*deck) != 52 && len(*emptyDeck) != 0 {
		t.Error("Card wasn't reinserted correctly")
	}
}

func TestDeckString(t *testing.T) {
	deck := NewDeck(1)

	fmt.Println(DeckString(deck))
}