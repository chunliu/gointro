package main

import "fmt"

func fib(out chan<- int, n int) { // HL
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		out <- a
	}
	close(out)
}

func main() {
	c := make(chan int)
	go fib(c, 10) // HL

	for x := range c {
		fmt.Println(x)
	}
}

// file end OMIT
