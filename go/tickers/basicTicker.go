package main

import (
	"fmt"
	"time"
)

func tickBasically2() {
	fmt.Println("Experimenting with buffered tickers")
	burstyLimiter := make(chan time.Time, 3)

	for range 3 {
		burstyLimiter <- time.Now()
	}

	fmt.Println("len of burstyLimiter:", len(burstyLimiter))

	go func() { // Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
		for t := range time.Tick(200 * time.Millisecond) {
			fmt.Println(t)
			burstyLimiter <- t
		}
	}()

	time.Sleep(2 * time.Second)
}

func tickBasically() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}
