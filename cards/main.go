package main

func main() {

	cards := newDeck()
	//
	// cards.print()
	// cards.saveToFile("these_cards")

	// newCards := newDeckFromFile("these_cards")
	// fmt.Println()
	cards.shuffle()
	cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println()
	// remainingCards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
