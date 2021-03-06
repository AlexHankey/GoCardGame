package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of deck
// Which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	// 4 card suits
	cardSuits := []string{
		"Spades",
		"Diamonds",
		"Clubs",
		"Hearts"}
	// 13 card values
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Creates a new receiver called d
// which references cards from main.go

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Converts a slice to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Converts the created string into a byte slice
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
