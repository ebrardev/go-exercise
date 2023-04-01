package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("give me a number")
		return
	}

	year, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}

	var leap bool

	if year%400 == 0 {
		leap = true

	} else if year%100 == 0 {
		leap = false
	} else if year%4 == 0 {
		leap = true
	}

	if leap {
		fmt.Printf("%d is a leap year\n", year)
	} else {
		fmt.Printf("%d is not a leap year\n", year)
	}

}
