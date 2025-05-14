// this is basically a copy of the adapter pattern in net/http between Handler, Handlerfunc and HandleFunc, just to make it a bit easier to think about

package main

import "fmt"

func main() {
	f := MultiplierFunc(MutipleFunc)
	fmt.Println(f.Multiply(2, 3)) // prints 6
}

// the analogue for Handler
type Multiplier interface {
	Multiply(a, b int) int
}

// the analogue for HandlerFunc
type MultiplierFunc func(a, b int) int

func (f MultiplierFunc) Multiply(a, b int) int {
	return f(a, b)
}

func MutipleFunc(a, b int) int {
	return a * b
}
