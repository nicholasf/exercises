package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	// queue <- "three"
}

func main2() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		fmt.Println("ok")
		// done <- true
	}()

	go func() {
		fmt.Println("Running")
		for {
			fmt.Println("Checking ..")
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
	// <-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
