package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("File IO")

	data := []byte("Hello Go")

	// Print out the ASCII code written to file
	for i, b := range data {
		fmt.Printf("Char: %c ASCII: %d\n", b, b)
		_ = i // ingores index if we don't need it
	}

	err := os.WriteFile("data.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
	}

	dataRead, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(dataRead))

	file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()

	file.WriteString("This is a test\n")
	file.WriteString("This is only a test\n")
	file.WriteString("We are appending data to this file\n\n")
	file.WriteString("Let's append one more line\n")
}
