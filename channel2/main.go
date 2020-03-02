package main

import (
	"fmt"
	"time"
)

func add(c chan int) {
	for i := 1; i < 200; i++ {
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
	//go read(chan1)
	time.Sleep(time.Second * 2)
	fmt.Println("main close")
}
