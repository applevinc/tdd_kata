package balancedbrackets

import (
	"math/rand"
)

var StringSizeArray = [4]int{2, 4, 6, 8}
var BracketArray = [2]string{"[", "]"}

func getStringLength() int {
	i := rand.Intn(4)
	return StringSizeArray[i]
}

func pickBracket() string {
	i := rand.Intn(2)
	return BracketArray[i]
}

func Generate() string {
	sl := getStringLength()
	var s string

	for i := 0; i < sl; i++ {
		b := pickBracket()
		s = s + b
	}

	return s
}
