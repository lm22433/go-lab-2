package main

import (
	"log"
	"os"
	"runtime/trace"
	"time"
)

func foo(channel chan string) {
	initMsg := "ping"
	println("Foo is sending:", initMsg)
	channel <- initMsg

	for {
		msg := <-channel
		println("Foo has received:", msg)
		println()

		newMsg := "ping"
		println("Foo is sending:", newMsg)
		channel <- newMsg
	}
}

func bar(channel chan string) {
	for {
		msg := <-channel
		println("Bar has received:", msg)

		newMsg := "pong"
		println("Bar is sending:", newMsg)
		channel <- newMsg
	}
}

func pingPong() {
	channel := make(chan string)
	go foo(channel)
	go bar(channel)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
