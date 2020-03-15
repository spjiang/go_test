package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode // 这个表示指向下一个节点
}

func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	// 空链表
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，不能删除")
		return head
	}
	// 如果只有一个结点
	if temp.next == head {
		temp.next = nil
		return head
	}
	// 将helper 定位到链表的最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	// 如果只有两个以上的结点
	flag := true
	for {
		// 说明已经循环完,找到最后一个，但是最后一个没有进行比较
		if temp.next == head {
			break
		}
		if temp.no == id {
			if temp == head {
				head = head.next
			}
			// 恭喜已经找到，我们可以直接删除
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
			break
		}
		temp = temp.next     // 移动比较
		helper = helper.next // 移动【一旦找到要删除的结点，helper】
	}
	// 这里还有比较一次
	if flag { // 如果flag为真，则我们上面没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Println("没有找到对应的Id")
		}
	}
	return head
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	// 判断是不是添加一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	// 定义一个临时值，帮忙，找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

// 输出这个环形的链表
func ListCircleLink(head *CatNode) {

	fmt.Println("环形链表的情况如下...")
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表...")
		return
	}

	for {
		fmt.Printf("猫的信息为=[id=%d name=%s]\n", temp.no, temp.name)
		fmt.Println(*temp)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {
	// 这里我们初始化一个环形链表的头结点
	head := &CatNode{}
	// 创建一个猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
		next: nil,
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom",
		next: nil,
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom",
		next: nil,
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	head = DelCatNode(head, 1)

	ListCircleLink(head)
}
