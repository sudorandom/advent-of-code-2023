package main

import (
	"testing"
)

func TestHandClassFromHand(t *testing.T) {
	if HandClassFromHand("A2T63") != HighCard {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("4854J") != ThreeOfKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("TJTT3") != FourOfAKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("69664") != ThreeOfKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("ATT9A") != TwoPair {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("69959") != ThreeOfKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("39666") != ThreeOfKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("JJA59") != ThreeOfKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("A7799") != TwoPair {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("T8JTA") != ThreeOfKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("3333J") != FiveOfAKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("43Q48") != OnePair {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("66266") != FourOfAKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("65TTT") != ThreeOfKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("TT222") != FullHouse {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("JAAQT") != ThreeOfKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("9Q959") != ThreeOfKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("27573") != OnePair {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("J2228") != FourOfAKind {
		t.Fatalf("unexpected class!")
	}
	if HandClassFromHand("AJ555") != FourOfAKind {
		t.Fatalf("unexpected class!")
	}
}
