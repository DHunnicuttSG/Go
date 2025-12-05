package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	part := numbers[0:4]
	//numbers = append(numbers, 88, 87, 84, 83, 82, 81)
	part = append(part, 10, 20, 30, 40)
	fmt.Println(numbers)

	fmt.Println(cap(part), len(part))
	fmt.Println(cap(numbers))

	fmt.Println(part)

	numbers[3] = 77
	part[0] = 99
	fmt.Println(part)
	fmt.Println(numbers)
	fmt.Println("numbers:", numbers)
	fmt.Printf("addr: %p\n", &numbers[0])
	fmt.Printf("addr: %p\n", &part[0])

}
