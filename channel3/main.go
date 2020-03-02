package main

import (
	"fmt"
	"time"
)

func add(c chan int) {
	for i := 1; i < 10; i++ {
		time.Sleep(time.Second)
		c <- i
		fmt.Println("i", i)
	}
	close(c)
}

func read(c chan int) {
	for {
		time.Sleep(time.Second)
		val, ok := <-c
		if !ok {
			break
		}
		fmt.Println(val)
	}
}

func main() {
	var chan1 chan int
	chan1 = make(chan int, 10)
	go add(chan1)
	for {
		v, ok := <-chan1
		if !ok {
			break
		}
		fmt.Println("v:", v)
	}
	fmt.Println("main close")
}
