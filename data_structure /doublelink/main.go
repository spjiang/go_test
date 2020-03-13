package main

import "fmt"

// 定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode // 这个表示指向上一个节点
	next     *HeroNode // 这个表示指向下一个节点
}

// 给链表插入一个结点
// 编写第一种插入方法，在单链表的最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 思路
	//1、先找到该链表的最后结点
	//2、创建一个辅助结点【跑龙套，帮忙】
	temp := head
	for {
		if temp.next == nil { // 表示找到最后
			break
		}
		temp = temp.next // 让temp不断的指向下一个结点
	}
	//3、将newHeroNode加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp.next
}

// 给链表插入一个结点
// 编写第2种插入方法，根据no编号从小到大插入...
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 思路
	//1、找到适当的结点
	//2、创建一个辅助结点【跑龙套，帮忙】
	var flag bool
	flag = true

	temp := head
	for {
		if temp.next == nil { // 表示找到最后
			break
		} else if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，已经存在no=", newHeroNode.no)
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}

}

func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil { // 表示找到最后
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}

	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("id 没有找到")
	}

}

// 显示链表的所有结点信息
func ListHeroNode(head *HeroNode) {
	// 1、创建一个辅助结点，
	temp := head
	// 先判断该链表是不是一个空链表
	if temp.next == nil {
		fmt.Println("空空如也..")
		return
	}
	// 2、变量这个链表
	for {
		fmt.Printf("[%d, %s, %s]===\n", temp.next.no, temp.next.name, temp.next.nickname)
		// 判断是否链表后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	// 1、创建一个头节点
	head := &HeroNode{}
	//2、创建一个新的节点
	head1 := &HeroNode{
		no:       1,
		name:     "宋江1",
		nickname: "及时雨",
	}
	head2 := &HeroNode{
		no:       3,
		name:     "宋江3",
		nickname: "及时雨",
	}
	head3 := &HeroNode{
		no:       2,
		name:     "宋江2",
		nickname: "及时雨",
	}
	// 3、加入
	InsertHeroNode2(head, head1)
	InsertHeroNode2(head, head2)
	InsertHeroNode2(head, head3)
	// 4、显示
	ListHeroNode(head)
}
