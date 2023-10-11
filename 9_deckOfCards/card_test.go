package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Club})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Queen, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Clubs
	// Nine of Diamonds
	// Queen of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Incorrect number of cards in a new deck.")
	}
}
