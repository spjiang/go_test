package main

import "fmt"

// 约瑟夫问题，丢手绢

type Boy struct {
	No   int
	Next *Boy // 指向下一个小孩的指针
}

// 编写一个函数，构成单向的环形链表
// num:表示小孩的个数
// *Boy：返回该环形链表的第一个小孩指针
func AddBoy(num int) *Boy {
	first := &Boy{} // 空结点
	curBoy := &Boy{}
	// 判断
	if num < 1 {
		fmt.Println("num值不对...")
		return first
	}
	// 循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		//分析构成环形链表，需要一个辅助指针【帮忙的】
		//1、因为第一个小孩比较特殊
		if i == 1 {
			first = boy // 不要动
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			// 相当于当前这个curBoy 向后移动到boy地址上
			curBoy = boy
			// 对这个移动后的curBoy next 指向 first
			// 因为first 是第一个，curBoy现在成为最后一个，最后一个next始终要指向第一个
			// 从而构成环形链表
			curBoy.Next = first
		}
	}
	return first
}

// 显示单向的环形链表【遍历】
func ShowBoy(first *Boy) {
	// 处理一下如果环形链表为空
	if first.Next == nil {
		fmt.Println("这是一个空的链表,没有小孩....")
		return
	}
	// 创建一个指针，帮助遍历
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

/**
设置编号为1，2，3...N的N个人围坐一圈，约定编号为K（1<=K<=n）的人从1开始报数，
数到m 的那个人出列，它的下一位从1开始重新报数，
数到m的那个人又出列，以此类推，直到所有人出列为止，由此产生一个出队编号的序列
*/
// 分析思路
// 1、编写一个函数，playGame(first *Boy,startNo int, countNum int)
// 2、最后我们使用一个算法，按照要求，在环形链表中留下最后一个人

func PlayGame(first *Boy, startNo int, countNum int) {
	// 1.空的链表我们单独处理
	if first.Next == nil {
		fmt.Println("没有小孩...")
		return
	}
	// 留一个，判断startNo<=小孩的总数

	// 2.需要定义辅助指针，帮助我们删除小孩
	tail := first

	//3.让tail执行环形链表最后一个小孩，这个非常重要
	// 因为tail，在删除小孩时会用到
	for {
		if tail.Next == first { // 说明tail到了最后一个小孩
			break
		}
		tail = tail.Next
	}
	//4.让first 移动到startNo[后面我们删除小孩，就以first为准]
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	// 5.开始数countNum，然后就删除first指向的小孩
	for {
		// 开始数countNum-1次
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.No)
		// 删除first执向的小孩
		first = first.Next
		tail.Next = first
		// 判断如果tail == first,圈中只有一个小孩
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩编号为%d 出圈 \n", first.No)
}

func main() {
	first := AddBoy(5)
	ShowBoy(first)
	PlayGame(first, 2, 3)
}
