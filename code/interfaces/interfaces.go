package main

import (
	"fmt"
	"math"
)

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

// interface start OMIT
type Abser interface {
	Abs() float64
}

func CalcAbs(a Abser) {
	fmt.Printf("The Abs of a is %f\n", a.Abs())
}

// interface end OMIT

func main() {
	f := MyFloat(-math.Sqrt2)
	v := Vertex{10, 10}
	CalcAbs(f)
	CalcAbs(&v)
}
