package main

func main() {
	// cards := newDeckFromFile("my_cards.txt")
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// cards.save("my_cards.txt")

	// fmt.Println(cards.toString())

	// hand, remainingCards := deal(cards, 5)

	// fmt.Println("===== Cards in hand ======")
	// hand.print()

	// fmt.Println("===== Cards remaining ====")
	// remainingCards.print()

	// fmt.Println("===== CARDS =====")
	// cards.print()
}
