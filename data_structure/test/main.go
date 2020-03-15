package main

import "fmt"

type Boy struct {
	No   int
	Next *Boy // 指向下一个小孩的指针
}

func main() {

	first := &Boy{}
	curBoy := &Boy{}

	boy := &Boy{
		No: 1,
	}

	fmt.Printf("first:%p,curBoy:%p,boy:%p \n", first, curBoy, boy)

	first = boy // 不要动
	fmt.Printf("first:%p,curBoy:%p,boy:%p \n", first, curBoy, boy)

	curBoy = boy

	fmt.Printf("first:%p,curBoy:%p,boy:%p\n", first, curBoy, boy)

	curBoy.Next = boy //

	fmt.Printf("first:%v,curBoy:%v,boy:%v\n", first, curBoy, boy)
}
