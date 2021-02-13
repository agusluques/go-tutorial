package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// create a new type of 'deck'
type deck []string

// create new function for the 'deck' type
func newDeck() deck {
	result := deck{}

	cardPalos := []string{"Espada", "Basto", "Oro", "Copa"}

	for number := 1; number < 13; number++ {
		for _, palo := range cardPalos {
			result = append(result, strconv.Itoa(number)+" de "+palo)
		}
	}

	return result
}

// shuffle a deck
func (d *deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })
}

// deal a hand of size 'handSize'
func (d *deck) deal(handSize int) deck {
	hand := (*d)[:handSize]
	*d = (*d)[handSize:]

	return hand
}

// print all cards in deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println("Card ", i+1, " - ", card)
	}
}

// converts a deck to a []byte to store in file system
func (d deck) toByteSlice() []byte {
	return []byte(strings.Join([]string(d), ","))
}

// save a deck to the file 'filename'
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toByteSlice(), 0666)
}

// converts a []byte to deck to read a file
func byteSliceToDeck(bs []byte) deck {
	return deck(strings.Split(string(bs), ","))
}

// read the file 'filename' and return the deck
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}

	return byteSliceToDeck(bs)
}
