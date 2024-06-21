package main

import "fmt"

func main() {
	cards := []string{"Ace Diamond", newCard()}
	cards = append(cards, "Six of speeads")
	fmt.Println(len(cards))

	for i := 0; i < len(cards); i++ {
			fmt.Println(i+1,cards[i])
	}
}


func newCard() string{
	return "Five of diamond"
}





