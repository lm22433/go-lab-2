package main

import (
	"fmt"
	"time"
)

func sendMessages(receiver chan string) {
	// Create a slice of some strings to send.
	messages := []string{
		"ping",
		"pong",
		"pinggg",
	}

	// Send the 3 messages to the receiver
	for _, m := range messages {
		fmt.Println("sendMessages is sending:", m)
		receiver <- m

	}
}

func main() {
	// Create a channel for sending and receiving strings.
	messages := make(chan string)

	// Start a new goroutine that will send some messages.
	go sendMessages(messages)

	// Receive the 3 messages sent by the goroutine.
	// Modified so that only 4 messages are received from the go function.
	// When receiving 4 messages, we get deadlock because all goroutines are asleep.
	// Will just endlessly wait for a response that is never sent.
	for i := 0; i < 4; i++ {
		// Wait 1s between each receive.
		time.Sleep(1 * time.Second)
		receivedMessage := <-messages
		fmt.Println("Main has received:", receivedMessage)
	}
}
