package functions

import (
	"fmt"
)

func select_fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}


func Select() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("ゴルーチン実行")
		for i := 0; i < 10; i++ {
			fmt.Println(i, "回目" , <-c)
		}
		quit <- 0
	}()
	fmt.Println("関数実行")
	select_fibonacci(c, quit)
}