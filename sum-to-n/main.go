package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("give me 2 numbers")
		return
	}
	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("wrong numbers")
		return
	}
	var sum int
	for i := min; i <= max; i++ {
		sum += i
		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}

	}
	fmt.Printf("=%d\n", sum)

}
