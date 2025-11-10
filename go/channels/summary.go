package main

import "fmt"

func summarise() {
	sendToUnbufferedChannelOps()

}

// 2. Receive from an unbuffered channel: The receiving goroutine pauses until another goroutine sends a value.

func receiveFromUnbufferedChannelOps() {
	fmt.Println("Receive from an unbuffered channel")

}

// 1. Send to an unbuffered channel: The sending goroutine pauses until another goroutine is ready to receive the value.
func sendToUnbufferedChannelOps() {
	fmt.Println("Send to an unbuffered channel")

	unbuffered := make(chan string)

	go func() {
		msg := <-unbuffered
		fmt.Println("received message", msg)
	}()

	unbuffered <- "hello" // this will block because there is no goroutine to receive it

	fmt.Println("the next line blocks unless we use select default")

	select {
	case v := <-unbuffered:
		fmt.Println("v:", v)
	default:
		fmt.Println("no value received")
	}

	go func() {
		v, ok := <-unbuffered

		fmt.Println("v:", v, " ok:", ok, " ok will be true")
	}()

	unbuffered <- "goodbye" // this will block until something receives from unbuffered, such as the above goroutine

	close(unbuffered)

	v, ok := <-unbuffered // this will not block because the channel is closed now

	fmt.Println("v:", v, " ok:", ok, " v is the zero value of string, ok will be false")
}
