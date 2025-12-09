# Question 1

Create a function called swap to swap the values of two integers using pointers

Here is a skeleton outline for your code

package main

import "fmt"

func swap() {
    //Your code goes here
}

// Do not change the code in the main function
func main() {
    x := 10
    y := 20

    swap(&x, &y)
    fmt.Println(x, y)
}

# Question 2
Create a struct called **Rectangle** and a method called **Area**
and the code that goes with it to calculate the area of the rectangle
Be sure that the method Area is a method of Rectangle

package main

import "fmt"

//Your code goes here


// Do not change the code in the main function
func main() {
    rect := Rectangle{Width: 10, Height: 5} 
    fmt.Println("Area:", rect.Area)
    
}

# Question 3
Write a function to count the words in the text that is passed to the function

package main

import (
	"fmt"
	"strings"
)

func countWords(s string) map[string]int {

}

func main() {
	s := "go is fun and go is fast"
	fmt.Println(countWords(s))
	t := "this is a test this is only a test"
	fmt.Println(countWords(t))
}
