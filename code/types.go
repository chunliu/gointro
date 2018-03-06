package main

import (
	"fmt"
)

type Point struct {
	x float64
	y float64
}

func typeInferecne() {
	var x int
	x = 42
	y := 24
	var z float32 = float32(x) // explicit conversion
	fmt.Println(x, y, z)
}
