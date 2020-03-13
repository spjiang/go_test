package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int    // 表示我们栈最大可以存放数个数
	Top    int    // 表示栈顶，因为栈顶固定，因此我们直接用top
	arr    [5]int // 数组模拟栈
}

func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

// 变量栈
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	// 判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	// 先取值，再this.Top --
	val = this.arr[this.Top]
	this.Top--
	return val, nil

}
func main() {
	stack := &Stack{
		MaxTop: 5,  // 表示最多存放5个数到栈中
		Top:    -1, // 当栈顶为-1，表示栈为空
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.List()

}
