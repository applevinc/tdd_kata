package balancedbrackets

import (
	"reflect"
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

func TestIsBalanced(t *testing.T) {
	t.Run("balanced bracket string", func(t *testing.T) {
		s := "[[]]"
		got := isBalanced(s)
		want := "OK"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("not balanced bracket string", func(t *testing.T) {
		s := "[["
		got := isBalanced(s)
		want := "FAIL"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

func TestSplitToHalves(t *testing.T) {
	t.Run("return first halve", func(t *testing.T) {
		s := []string{"a", "b", "l", "e"}
		got, _ := firstHalf(s)
		want := "ab"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("return second halve", func(t *testing.T) {
		s := []string{"a", "b", "l", "e"}
		got, _ := secondHalf(s)
		want := "le"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("first half: return error for odd slice length", func(t *testing.T) {
		s := []string{"v", "i", "n", "c", "e", "n", "t"}
		_, got := firstHalf(s)
		err := OddSliceLength

		if got != err {
			t.Errorf("got %q want %q", got, err)
		}
	})

	t.Run("second half: return error for odd slice length", func(t *testing.T) {
		s := []string{"v", "i", "n", "c", "e"}
		_, got := secondHalf(s)
		err := OddSliceLength

		if got != err {
			t.Errorf("got %q want %q", got, err)
		}
	})

	t.Run("return [dave] as [da ve]", func(t *testing.T) {
		s := "dave"
		got := splitToHalves(s)
		want := []string{"da", "ve"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
