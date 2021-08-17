package balancedbrackets

import (
	"math/rand"
	"strings"
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

func firstHalf(strArr []string) string {
	var result string

	for i := 0; i < len(strArr)/2; i++ {
		result = result + strArr[i]
	}
	return result
}

func secondHalf(strArr []string) string {
	var result string
	startIndex := len(strArr) / 2

	for i := startIndex; i < len(strArr); i++ {
		result = result + strArr[i]
	}
	return result
}

func splitToHalves(s string) []string {
	result := []string{}
	strArr := strings.Split(s, "")
	result = append(result, firstHalf(strArr), secondHalf(strArr))
	return result
}

func isBalanced(s string) string {
	if s != "[]" {
		return "FAIL"
	}
	return "OK"
}
