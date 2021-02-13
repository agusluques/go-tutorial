package main

import "fmt"

func main() {
	// var card string = "Ancho de espada"
	// card := "Ancho de espada"
	// card2 := newCard()
	// cards := []string{newCard(), "7 de espada"}
	// cards = append(cards, "7 de oro")

	// fmt.Println(card)
	// fmt.Println(card2)
	// for _, card := range cards {
	// 	fmt.Println(card)
	// }

	myDeck := newDeck()

	fmt.Println("My initial deck")

	myDeck.shuffle()
	myDeck.print()
}

// func newCard() string {
// 	return "Ancho de basto"
// }
