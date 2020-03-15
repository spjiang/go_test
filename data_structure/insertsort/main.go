package main

import "fmt"

func InsertSort(arr *[5]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数组后移动
			insertIndex--
		}
		// 已经找到插入位置，insertIndex 对应值比insertVal大，所以只能插入到insertIndex+1 这个位置
		if insertIndex+1 != i { // 本身这个insertVal值是在原来的位置
			arr[insertIndex+1] = insertVal
		}
	}
}

func main() {
	// 定义一个数组，从大到小
	arr := [5]int{23, 0, 12, 56, 34}
	InsertSort(&arr)
	fmt.Println(arr)
}
