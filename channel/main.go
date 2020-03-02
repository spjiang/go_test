package main

import (
	"fmt"
	"time"
)

func putNum(putChan chan int) {
	fmt.Println("插入管道开始...")
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Millisecond * 200)
		fmt.Println("插入管道:", i)
		putChan <- i
	}
	fmt.Println("插入管道退出...")
	close(putChan)
}

func calNum(putChan chan int, calChan chan int, resChan chan bool) {
	var flag bool
	for {
		num, ok := <-putChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			// 不是素数
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			calChan <- num
		}
	}
	fmt.Println("有一个calNum协程因为却不到数据退出....")
	resChan <- true

}

func main() {
	fmt.Println("main start")
	putChan := make(chan int, 5)
	calChan := make(chan int, 2)
	resChan := make(chan bool, 4)
	// 写入数据
	go putNum(putChan)
	// 4个协程读出数据
	for i := 0; i < 4; i++ {
		go calNum(putChan, calChan, resChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-resChan
		}
		close(calChan)
	}()

	/*go func() {
		for {
			res, ok := <-calChan
			if !ok {
				break
			}
			fmt.Println("素数：", res)
		}
	}()*/

	for {
		res, ok := <-calChan
		if !ok {
			break
		}
		fmt.Println("素数：", res)
	}

	fmt.Println("main线程退出")

}
