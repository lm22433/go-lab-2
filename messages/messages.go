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
	// Modify the make func call so that the messages channel
	// is a buffered channel of size 3.
	// This changes the behaviour of the program because there is space
	// in the buffer for all messages to be sent amd fill the buffer concurrently.
	// As a result, instead of having it sent and received in order, they are all
	// sent at once and received at once.
	messages := make(chan string, 3)

	// Start a new goroutine that will send some messages.
	go sendMessages(messages)

	// Receive the 3 messages sent by the goroutine.
	// When modified to receive 4 messages, we get deadlock because all goroutines are asleep.
	// Will just endlessly wait for a response that is never sent.
	for i := 0; i < 3; i++ {
		// Wait 1s between each receive.
		time.Sleep(1 * time.Second)
		receivedMessage := <-messages
		fmt.Println("Main has received:", receivedMessage)
	}
}
