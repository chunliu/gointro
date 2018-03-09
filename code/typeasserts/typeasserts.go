package main

import (
	"fmt"
	"io"
	"os"
)

func do(i interface{}) {
	switch v := i.(type) { // Type switch // HL
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var w io.ReadWriter
	w = os.Stdout
	if f, ok := w.(*os.File); ok { // Type assertion // HL
		do(f)
	}
	do(21)
	do("hello")
	do(true)
}

//do OMIT
