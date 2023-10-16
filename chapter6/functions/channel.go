package functions

import (
	"fmt"
)

func sum(s []int, c chan int) {
	// fmt.Println("sum", s)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum

	c <- 1
}

func Channel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 4)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	fmt.Println(c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}