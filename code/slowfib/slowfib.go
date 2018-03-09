package main // HL

import ( // HL
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
func main() { // HL
	fibN := make(chan int)
	const n = 42
	go getFib(n, fibN) // HL
	fmt.Printf("Calculation starts ... ")
	fmt.Printf("\rFibonacci(%d) = %d\n", n, <-fibN)
}
