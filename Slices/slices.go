package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	part := numbers[0:4]

	fmt.Println(part)

	numbers[4] = 77
	part[0] = 99
	fmt.Println(part)

	fmt.Println("numbers:", numbers)
	fmt.Printf("addr: %p\n", &numbers)
	fmt.Printf("addr: %p\n", &part)

}
