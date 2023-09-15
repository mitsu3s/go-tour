package functions

import "fmt"

var check bool

func AddInt(x, y int) int {
	fmt.Println(check)
	return x + y
}

func SwapString(x, y string) (string, string) {
	return y, x
}
