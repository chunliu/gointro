package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() (float64, error) {
	// if v == nil {
	// 	return 0, fmt.Errorf("parameter cannot be nil")
	// }
	return math.Sqrt(v.X*v.X + v.Y*v.Y), nil
}

func main() {
	var v *Vertex
	f, err := v.Abs()
	if err != nil { // basic error handling // HL
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("The Abs of a is %f\n", f)
}

// code end OMIT
