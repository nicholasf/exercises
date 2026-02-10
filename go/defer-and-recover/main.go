package main

import "fmt"

func main() {
	defer fmt.Println("This will print last")
	fmt.Println("This will print first")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("Something went wrong!")
}
