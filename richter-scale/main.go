package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me a magnitude pf the earthquake")
		return
	}
	richter, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("I can't understand the magnitude of the earthquake")
		return
	}

	var desc string

	switch r := richter; {
	case r < 2:
		desc = "Micro"
	case r >= 2 && r < 3:
		desc = "Very minor"
	case r >= 3 && r < 4:
		desc = "Minor"
	case r >= 4 && r < 5:
		desc = "Light"
	case r >= 5 && r < 6:
		desc = "Moderate"
	case r >= 6 && r < 7:
		desc = "Strong"
	case r >= 7 && r < 8:
		desc = "Major"
	case r >= 8 && r < 10:
		desc = "Great"

	default:
		desc = "Meteoric"

	}

	fmt.Printf("%.2f is %s\n", richter, desc)

}
