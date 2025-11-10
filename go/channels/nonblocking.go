package main

import "fmt"

func nonBlockingChannelOps() {
	messages := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	// <-messages
}
