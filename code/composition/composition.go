package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}
type Circle struct {
	Center Point // try embedded
	Radius int
}
type Wheel struct {
	Circle Circle // try embedded
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v", w)
}

// main end OMIT
