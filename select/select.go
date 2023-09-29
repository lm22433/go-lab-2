package main

import (
	"fmt"
	"time"
)

// slowSender sends a string every 2 seconds.
func slowSender(c chan<- string) {
	for {
		time.Sleep(2 * time.Second)
		c <- "I am the slowSender"
	}
}

// fastSender sends consecutive ints every 500 ms.
func fastSender(c chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		c <- i
	}
}

func fasterSender(c chan<- []int) {
	slice := []int{1, 2, 3}
	for {
		time.Sleep(200 * time.Millisecond)
		c <- slice
	}
}

// main starts the two senders and then goes into an infinite loop of receiving their messages.
func main() {
	slices := make(chan []int)
	go fasterSender(slices)
	ints := make(chan int)
	go fastSender(ints)
	strings := make(chan string)
	go slowSender(strings)

	for { // = while(true)
		// We select which print to chose based on which channel has a value
		// The slower sender takes longer to populate the channel
		// the faster sender takes less time
		// so we have more int messages than string messages
		select {
		case s := <-strings:
			fmt.Println("Received a string", s)
		case i := <-ints:
			fmt.Println("Received an int", i)
		case xs := <-slices:
			fmt.Println("Received a slice", xs)
		default:
			fmt.Println("--- Nothing to receive, sleeping for 3s...")
			time.Sleep(3 * time.Second)
		}
	}
}
