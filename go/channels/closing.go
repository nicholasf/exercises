package main

import (
	"fmt"
	"time"
)

func closing() {
	ch := make(chan int)

	select {
	case v, more := <-ch:
		fmt.Println("more? ", more, "value:", v, " at time:", time.Now())
	default:
		fmt.Println("1. No value received, no indication of channel state at time:", time.Now())
	}

	close(ch)

	select {
	case v, more := <-ch:
		fmt.Println("2. more? ", more, "value:", v, " at time:", time.Now())
	default:
		fmt.Println("No value received, no indication of channel state at time:", time.Now())
	}

}
