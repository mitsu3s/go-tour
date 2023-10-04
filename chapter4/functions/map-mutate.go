package functions

import (
	"fmt"
)

func MapMutate() {
	m := make(map[string]int)

	m["math"] = 90
	m["science"] = 80
	m["english"] = 70
	fmt.Println("The value of m[\"math\"] is", m["math"])

	delete(m, "math")
	fmt.Println("The value of m[\"math\"] is", m["math"])

	v, ok := m["math"]
	fmt.Println("The value of v is", v, "and ok is", ok)
}
