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

	switch r := m; {
	case r >= 10:
		fmt.Printf("%.2f is massive", m)
	case r >= 8:
		fmt.Printf("%.2f is great", m)
	case r >= 7:
		fmt.Printf("%.2f is major", m)
	case r >= 6:
		fmt.Printf("%.2f is strong", m)
	case r >= 5:
		fmt.Printf("%.2f is moderate", m)
	case r >= 4:
		fmt.Printf("%.2f is light", m)
	case r >= 3:
		fmt.Printf("%.2f is minor", m)
	case r >= 2:
		fmt.Printf("%.2f is very minor", m)
	case r >= 0:
		fmt.Printf("%.2f is micro", m)
	}
}
