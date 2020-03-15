package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

// 前序遍历【先输出root结点，然后在输出左子树，在输出右子树】
func PreOrder(node *Hero) {
	// zzzsss
	if node != nil {
		fmt.Printf("no=%d, name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历【，先输出左子树，然后在输出root结点，在输出右子树】
func InfixOrder(node *Hero) {
	if node != nil {
		PreOrder(node.Left)
		fmt.Printf("no=%d, name=%s\n", node.No, node.Name)
		PreOrder(node.Right)
	}
}

// 后序遍历【，先输出左子树，在输出右子树，然后在输出root结点】
func PostOrder(node *Hero) {
	if node != nil {
		PreOrder(node.Left)
		PreOrder(node.Right)
		fmt.Printf("no=%d, name=%s\n", node.No, node.Name)
	}
}

func main() {
	// 构建一个二叉树
	root := &Hero{
		No:   1,
		Name: "宋江",
	}
	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}

	node10 := &Hero{
		No:   10,
		Name: "tom",
	}
	node12 := &Hero{
		No:   12,
		Name: "jack",
	}
	left1.Left = node10
	left1.Right = node12

	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}
	root.Left = left1
	root.Right = right1

	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}
	right1.Right = right2

	PreOrder(root)

}
