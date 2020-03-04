package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int //数组=>模拟队列
	front   int    //表示指向列首
	rear    int    // 表示指向队列尾部
}

// 添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	// 重要的提示，rear是队列尾部（含最后一个元素）
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}
	this.rear++ // rear 后移动
	this.array[this.rear] = val
	return
}
func (this *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if this.front == this.rear {
		err = errors.New("queue empty")
		return
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

// 显示队列，找到队列首部，然后到遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	// this.front 不包含队首元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t\n", i, this.array[i])
	}
}

func main() {
	// 先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从对了获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示退出队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			fmt.Println("get")
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("从队列中取出了一个数=%d\n", val)
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
