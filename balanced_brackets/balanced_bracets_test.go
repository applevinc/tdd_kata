package balancedbrackets

import (
	"testing"
)

func TestGetStringSize(t *testing.T) {
	var prev int

	for i := 0; i < 10; i++ {
		got := getStringLength()

		if got == prev {
			t.Fatalf("returned the same int")
		}
		prev = got
	}
}

func TestPickBracket(t *testing.T) {

	var prev string

	for i := 0; i < 10; i++ {
		got := pickBracket()

		if got == prev {
			t.Fatalf("returned the same string")
		}
		prev = got
	}
}

func TestGenerate(t *testing.T) {

	var prev string

	for i := 0; i < 10; i++ {
		got := Generate()

		if got == prev {
			t.Fatalf("returned the same string")
		}
		prev = got
	}
}
