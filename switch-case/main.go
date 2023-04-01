package main

import (
	"fmt"
	"os"
)

func main() {

	city := os.Args[1]

	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Rome", "Milan":
		fmt.Println("Italy")
	case "Madrid", "Barcelona":
		fmt.Println("Spain")
	default:
		fmt.Printf("I don't know where %s is.\n", city)
	}
}
