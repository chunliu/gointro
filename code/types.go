package main

import (
	"fmt"
)

func typeInferecne() {
	var x int
	x = 42
	y := 24
	var z float32 = float32(x) // explicit conversion
	fmt.Println(x, y, z)
}
