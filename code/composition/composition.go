package main

import (
	"fmt"
)

// composed start OMIT
type Point struct {
	X, Y int
}
type Circle struct {
	Point  // embedded. equivalent to Center Point
	Radius int
}
type Wheel struct {
	Circle // embedded. equivalent to Circle Circle
	Spokes int
}

// composed end OMIT

func main() {
	// f := MyFloat(-math.Sqrt2)
	// fmt.Printf("The Abs of %f is %f\n", -math.Sqrt2, f.Abs())
	// v := &Vertex{10, 10}
	// fmt.Printf("The Abs of %+v is %f", v, v.Abs())

	// main start OMIT
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v", w)
	// main end OMIT
}
