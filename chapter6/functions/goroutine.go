package functions

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func count(n int, s string) {
	fmt.Println(s)
	for i := 0; i < n; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	}
}

func Goroutine() {
	go say("world")
	say("hello")

	go count(5 ,"first")
	count(10, "second")
}