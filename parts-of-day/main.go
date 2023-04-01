package main

import (
	"fmt"

	"time"
)

func main() {

	h := time.Now().Hour()
	fmt.Println("current hour is", h)

	switch {
	case h >= 18:
		fmt.Println("Good evening")
	case h >= 12:
		fmt.Println("Good afternoon")
	case h >= 6:
		fmt.Println("Good morning")
	default:
		fmt.Println("Good night")
	}

}
