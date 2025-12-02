package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func getValidAge(reader *bufio.Reader) int {
	for {
		ageStr := getInput("Enter your age: ", reader)
		age, err := strconv.Atoi(ageStr)

		if err == nil && age > 0 {
			return age
		}

		fmt.Println("Please enter a valid positive number for age.")
	}
}

func main() {
	var name string
	var age int
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)

	// fmt.Print("Enter your age: ")
	// fmt.Scanln(&age)

	name = getInput("Enter your name: ", reader)
	age = getValidAge(reader)

	fmt.Println("\n---- User Info ----")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
}
