package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int) // HL
	squares := make(chan int)  // HL

	go func() { // HL
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() { // HL
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Printf("%d\n", x)
	}
}
