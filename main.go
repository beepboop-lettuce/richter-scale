package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}
	userInput := os.Args[1]
	m, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	switch m; {
	case m >= 10:
		fmt.Printf("%.2f is massive", m)
	case m >= 8:
		fmt.Printf("%.2f is great", m)
	case m >= 7:
		fmt.Printf("%.2f is major", m)
	case m >= 6:
		fmt.Printf("%.2f is strong", m)
	case m >= 5:
		fmt.Printf("%.2f is moderate", m)
	case m >= 4:
		fmt.Printf("%.2f is light", m)
	case m >= 3:
		fmt.Printf("%.2f is minor", m)
	case m >= 2:
		fmt.Printf("%.2f is very minor", m)
	case m >= 0:
		fmt.Printf("%.2f is micro", m)
	}
}
