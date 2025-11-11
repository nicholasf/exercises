package main

import "fmt"

func main() {
	fmt.Println("Running main")
	// nonBlockingChannelOps()
	// closing()
	// ranger()
	// summarise()
	waitGroup()
}

/*

1. Send to an unbuffered channel: The sending goroutine pauses until another goroutine is ready to receive the value.

2. Receive from an unbuffered channel: The receiving goroutine pauses until another goroutine sends a value.

3. Send to a buffered channel: The sending goroutine pauses only if the buffer is full. If there is space, the value is placed in the buffer and the send operation completes without waiting for a receiver.

4. Receive from a buffered channel: The receiving goroutine pauses only if the buffer is empty.

5. Also practice for .. range over channels, which will just break if the channel is closed.
*/
