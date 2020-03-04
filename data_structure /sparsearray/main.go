package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// 1、先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //墨子
	chessMap[2][3] = 2 //蓝子
	// 2、输出看看原始的数组
	for i, v := range chessMap {
		fmt.Printf("行：%d ++++", i)
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 3、转成稀疏数组
	// 思路
	// （1）、遍历chessMap，如果我们发现有一个元素的值不为0，创建一个node结构体
	// （2）、将其放入到对应的切片即可
	var sparseArray []ValNode

	// 标准的一个稀疏数组应该还有一个 记录元素的二维数组的规模（行，列，默认值）
	// 创建一个ValNode值节点
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArray = append(sparseArray, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArray = append(sparseArray, valNode)
			}
		}
	}
	// 输出稀疏数组
	fmt.Println("当前的稀疏数组是：：：")
	for i, valNode := range sparseArray {
		fmt.Printf("%d：%d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	// 将这个稀疏数组，存盘chessMap.data
	// 如何恢复原始的数据

	//1、打开的这chessMap.data=>恢复原始数组
	//2、使用这个稀疏数组恢复

	fmt.Println("恢复后的数据...")
	var chessMap2 [11][11]int
	for i, valNode := range sparseArray {
		if i == 0 {
			continue
		}
		chessMap2[valNode.row][valNode.col] = valNode.val
	}
	// 打印恢复后的数据
	for i, v := range chessMap2 {
		fmt.Printf("行：%d ++++", i)
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
