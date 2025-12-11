package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printMessage("Hello from Goroutine!")
	printMessage("Hello from Main!")
}

// Notice how both messages interleave. That’s concurrency in action!
