package functions

import (
	"fmt"
	"math"
)

func mySqrt(x float64) string {
	if x < 0 {
		return mySqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func myPow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func If() {
	fmt.Println("mySqrt")
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(-4))
	fmt.Println("\nmyPow")
	fmt.Println(myPow(3, 2, 10))
	fmt.Println(myPow(3, 3, 20))
}
