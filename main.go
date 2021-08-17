package main

import (
	"fmt"
	"math/rand"
	"time"

	bk "tdd_kata/balanced_brackets"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(bk.Generate())
}
