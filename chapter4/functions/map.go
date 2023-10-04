package functions

import (
	"fmt"
)

type MapType struct {
	Latitude, Longitude float64
}

func Map() {
	var m map[string]MapType

	m = make(map[string]MapType)
	m["Bell Labs"] = MapType{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
