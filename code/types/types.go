package main

import (
	"fmt"
)

type Point struct {
	x float64
	y float64
}

func main() {
	// START OMIT
	var x int
	x = 42
	y, s := 24, "say hi"
	var z float32 = float32(x) // explicit conversion
	fmt.Println(x, y, z, s)
	// END OMIT
}
