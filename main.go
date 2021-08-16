package main

import (
	"fmt"
	"math/rand"
	bk "tdd_kata/balanced_brackets"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(bk.Generate())
}
