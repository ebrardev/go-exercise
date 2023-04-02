package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	t := time.Now()
	rand.Seed(t.UnixNano())
	rand.Seed(100)
	guess := 13

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d\n", n)
	}
	fmt.Println()
}
