package main

import "fmt"

//func main() {
//	c := make(chan os.Signal, 0)
//	signal.Notify(c)
//
//	// Block until a signal is received.
//	s := <-c
//	fmt.Println("Got signal:", s) //Got signal: terminated
//
//}

func main() {
	var c = make(chan int)
	var a string

	go func() {
		a = "hello world"
		<-c
	}()

	c <- 0
	fmt.Println(a)
}
