package main

import (
	"fmt"
	"go-tour/chapter2/functions"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(functions.AddInt(42, 13))

	a, b := functions.SwapString("hello", "world")
	fmt.Println(a, b)
}
