package main

import (
	"fmt"
	"os"
)

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (this *Emp) ShowMe() {
	fmt.Printf("链表Id：%d，雇员Id:%d,名称：%s", this.Id%7, this.Id, this.Name)
}

// 定义一个EmpLink
// 我们这里的EmpLink 不带表头，即第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

// 添加员工方法，保证添加时，编号从小到大
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head   // 这是一个辅助指针
	var pre *Emp = nil // 这是一个辅助指针pre在cur前面
	// 如果当前的EmpLink就是一个空链表
	if cur == nil {
		this.Head = emp
		return
	}
	// 如果不是一个空链表,给emp找到对应的位置插入
	// 思路是让cur 和emp比较，然后让pre保持在cur前面
	for {
		if cur != nil {
			// 比较
			if cur.Id > emp.Id {
				// 找到位置break
				break
			}
			pre = cur // 保持同步
			cur = cur.Next
		} else {
			break
		}
	}
	// 退出时，我们看下是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur

}
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表 %d  为空\n", no)
		return
	}
	cur := this.Head
	fmt.Printf("链表 %d:", no)
	for {
		if cur != nil {
			fmt.Printf("雇员Id:%d,名字:%s",  cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
	return
}

// 定义一个hashtable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

func (this *HashTable) HashFun(id int) int {
	return id % 7 // 得到一个值，就是对于的链表下标
}

func (this *HashTable) FindById(id int) *Emp {
	// 使用散列函数，确定到那个链表中
	linkNo := this.HashFun(id)
	// 使用对应的链表添加
	return this.LinkArr[linkNo].FindById(id)
}

// 给HashTable 编写一个Insert 雇员的方法
func (this *HashTable) Insert(emp *Emp) {
	// 使用散列函数，确定添加到那个链表中
	linkNo := this.HashFun(emp.Id)
	// 使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for {
		fmt.Println("==========雇员系统菜单============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("exit 退出系统")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员ID")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员名称")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.ShowAll()
		case "find":
			fmt.Println("请输入雇员ID")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Println("没有找到雇员ID：", id)
			} else {
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("命令输入错误，请重新输入")
		}

	}
}
