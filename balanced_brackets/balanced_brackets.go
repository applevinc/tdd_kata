package balancedbrackets

import (
	"errors"
	"log"
	"math/rand"
	"strings"
)

const (
	OK   = "OK"
	FAIL = "FAIL"
)

var OddSliceLength = errors.New("odd slice length")

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

func firstHalf(strArr []string) (string, error) {
	if !isEven(len(strArr)) {
		return "", OddSliceLength
	}

	var result string
	for i := 0; i < len(strArr)/2; i++ {
		result = result + strArr[i]
	}
	return result, nil
}

func secondHalf(strArr []string) (string, error) {
	if !isEven(len(strArr)) {
		return "", OddSliceLength
	}

	var result string
	startIndex := len(strArr) / 2
	for i := startIndex; i < len(strArr); i++ {
		result = result + strArr[i]
	}
	return result, nil
}

func splitToHalves(s string) []string {
	result := []string{}
	strArr := strings.Split(s, "")
	h1, _ := firstHalf(strArr)
	h2, _ := secondHalf(strArr)
	result = append(result, h1, h2)
	return result
}

func isBalanced(s string) string {
	h := splitToHalves(s)
	if len(h) > 2 {
		log.Fatalf("splitted string slice - %v elements is more than 2", h)
	}

	h1 := h[0]
	h2 := h[1]
	if h1 == h2 {
		return FAIL
	}

	if s != "[]" {
		return FAIL
	}
	return OK
}

func isEven(n int) bool {
	return n%2 == 0
}
