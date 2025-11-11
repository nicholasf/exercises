package main

import "fmt"

func ranger() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	k, ok := <-queue

	fmt.Println("k:", k, " ok:", ok)

	<-queue

	fmt.Println("k:", k)
}
