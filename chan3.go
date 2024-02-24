package main

import (
	"fmt"
	"time"
)

func mult(ch chan int, n int, m int) {
	ch <- n * m
}

func main() {

	ch := make(chan int, 2)
	c2 := make(chan int, 2)

	go mult(ch, 2, 3)
	go mult(ch, 3, 4)
	go mult(c2, 5, 6)
	time.Sleep(1 * time.Second)

	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case m := <-c2:
			fmt.Println(m)
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("default")

		}
	}

	//e := <-ch
	//fmt.Println(e)
}
