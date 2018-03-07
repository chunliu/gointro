package main

import (
	"fmt"
	"math"
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

// method start OMIT
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// method end OMIT

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
