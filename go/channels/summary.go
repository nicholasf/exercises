package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func summarise() {
	// sendToUnbufferedChannelOps()
	// receiveFromUnbufferedChannelOps()
	// sendToBufferedChannelOps()
	receiveFromBufferedChannelOps()
}

// 4. Receive from a buffered channel: The receiving goroutine pauses only if the buffer is empty.
func receiveFromBufferedChannelOps() {
	ns := make(chan int, 20)

	go func(numbers chan int) {
		for {
			n := <-numbers
			fmt.Println("received number", n)
		}
	}(ns)

	for i := 0; i < 10; i++ {
		fmt.Println("sending number", i)
		ns <- i
	}

	time.Sleep(1 * time.Second)

	for i := 10; i < 20; i++ {
		fmt.Println("sending number", i)
		ns <- i
	}

	// if the main goroutine exits before the receiving goroutine has a chance to process all the numbers, we won't see all the output
	time.Sleep(1 * time.Second)
}

// 3. Send to a buffered channel: The sending goroutine pauses only if the buffer is full.
// If there is space, the value is placed in the buffer and the send operation completes without waiting for a receiver.

func sendToBufferedChannelOps() {
	fmt.Println("Send to a buffered channel")

	buffered := make(chan string, 5)

	go func() {
		s := "this is a long sentence of more than the buffer size of buffered"
		bits := strings.Split(s, " ")

		for _, b := range bits {
			buffered <- b
			fmt.Println("sent string", b)
		}
		close(buffered)
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("Ranging over buffered channel")

	for {
		i, ok := <-buffered
		fmt.Println("received string", i)

		if !ok {
			break
		}
	}
}

func sendToBufferedChannelOps2() {
	fmt.Println("Send to a buffered channel")

	buffered := make(chan string, 5)

	go func() {
		for i := range buffered {
			fmt.Println("received message", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 20; i++ {
		buffered <- strconv.Itoa(i)

		fmt.Println("sent message", i)
		if i%2 == 0 {
			time.Sleep(1 * time.Second)
		}
	}

}

// 2. Receive from an unbuffered channel: The receiving goroutine pauses until another goroutine sends a value.
func receiveFromUnbufferedChannelOps() {
	fmt.Println("Receive from an unbuffered channel")

	unbuffered := make(chan string)

	go func() {
		for {
			fmt.Println("Checking...")
			if msg, ok := <-unbuffered; ok {
				fmt.Println("received message", msg)
			}

		}
	}()

	unbuffered <- "hello"
	time.Sleep(2 * time.Second)
	unbuffered <- "goodbye"
	time.Sleep(2 * time.Second)

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
