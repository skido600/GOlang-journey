package main

import (
	"fmt"
	"os"
	"time"
)

// What is a Goroutine?
// A goroutine is like a lightweight thread.
// It allows Go to run multiple tasks at the same time.
// Created with the keyword go.

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		fmt.Println(time.Now())
	}
}
func main() {
	// go printMessage("Hello") // runs in background
	// printMessage("World")

	// messages := make(chan string)

	// go func() { messages <- "ping" }()

	// msg := <-messages
	// fmt.Println(msg)

	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(data))
}
