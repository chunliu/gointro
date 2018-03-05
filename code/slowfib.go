package main

import (
	"fmt"
)

func getFib(n int, out chan int) {
	out <- fib(n)
}
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func main() {
	fibN := make(chan int)
	fmt.Printf("Start")
	const n = 45
	go getFib(n, fibN)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, <-fibN)
}
