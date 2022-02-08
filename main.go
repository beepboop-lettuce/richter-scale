package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
	}

	userInput := os.Args[1]

	switch m, _ := strconv.ParseFloat(userInput, 64); {
	case m >= 10:
		fmt.Printf("%.2f is massive")
	case m >= 8:
		fmt.Printf("%.2f is great")

	}
}
