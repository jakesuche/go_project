package main

import "fmt"

// create a new type of deck
// which is a slice of strings
// adorenke


type deck []string



func newDeck() deck {
	cards := deck{}
	cardsSuites := []string {"Spades", "Diamonds", "Hearts", "Clubs", "Kings"}
	cardValues := []string {"Ace", "Two", "Three", "Four", "Five"}
	for _, suit := range cardsSuites{
		for _, values := range cardValues {
			cards = append(cards, suit+ " of " + values)
		}
	}

	return cards
	
}

func (d deck) print(){	
	for i, card := range d {
		fmt.Println(i, card)
	}
}



func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}