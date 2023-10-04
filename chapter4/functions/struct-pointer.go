package functions

import (
	"fmt"
)

type Coordinate struct {
	X int
	Y int
}

func StructPointer() {
	v := Coordinate{1, 2}
	p := &v
	p.X = 4
	fmt.Println(v)

	v1 := Coordinate{1, 2}
	v2 := Coordinate{X: 1}
	v3 := Coordinate{}
	p = &Coordinate{1, 2}

	fmt.Println(v1, p, v2, v3)
}
