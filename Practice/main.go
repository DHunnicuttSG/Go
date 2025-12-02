package main

import (
	"fmt"
)

func add(s []int) []int {
	s = append(s, 100) // might create new array
	return s
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero!")
	}
	return a / b, nil
}

func showSlice() {
	s := []int{1, 2, 3}

	s = make([]int, 3, 5)

	s = append(s, 99)
	s = append(s, 98)
	s = append(s, 100)
	fmt.Println(s)

	fmt.Println(len(s))
	fmt.Println(cap(s))

	nums := []int{1, 2, 3}
	add(nums)
	fmt.Println(nums) // might NOT have 100
}

func getName() (name string) {
	name = "liz"
	return
}

func main() {
	fmt.Println("Practice")

	//showSlice()

	fmt.Println(divide(9, 7))
	fmt.Println(divide(10, 0))

	fmt.Println(getName())

}
