package main

import "github.com/cs3305group13/project/Cards"

func main() {
	deck := Cards.CreateDeck()
	Cards.Shuffle(deck)
	Cards.Deal(deck, 2)
	Cards.Shuffle(deck)
	Cards.Deal(deck, 2)
}
