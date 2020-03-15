package main

import "fmt"

func SelectSort(arr *[5]int) {
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		for i := j; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		// 交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}
}

func main() {
	// 定义一个数组，从大到小
	arr := [5]int{10, 34, 19, 100, 30}
	SelectSort(&arr)
	fmt.Println(arr)

}
