package main




func main() {
	cards := newDeck()
	hand, remain := deal(cards,5)
	hand.print()
	remain.print()

}






