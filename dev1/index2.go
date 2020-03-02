package main

import (
	"fmt"

)

type ABC interface {
	say()
}

type ABB struct {

}

func (c *ABB) say()  {
	fmt.Println("ABC")
}

func main() {
	var a point
	var abb ABB = ABB{}
	var abc ABC
	abc = &abb
	abc.say()


	// 定义一个map切片
	m := []map[string]string{
		{"ccc": "ccc"},
		{"bbb": "bbb"},
	}
	fmt.Println(m)

	// x(3)
	var arr [3]int = [3]int{1, 2, 3}
	var arr2 [3]int = [3]int{4, 5, 6}
	//var arr3 [3]int = [3]int{4, 5, 6}

	var slice []int = arr[:]
	var slice2 []int = arr2[:]

	var slice3 []*int = make([]*int, 10)
	var int1 int = 5
	var int2 int = 5
	var int3 int = 5

	slice3[0] = &int1
	slice3[1] = &int2
	slice3[2] = &int3

	fmt.Println("slice3:", slice3)
	fmt.Println("slice3 value:", *slice3[0])

	type Test struct {
		Name string
		Age  int
	}
	var test1 *Test = new(Test)
	(*test1).Name = "jiang"
	(*test1).Age = 10

	fmt.Println(*test1)

	slice4 := make([]*Test, 10)
	slice4[0] = test1

	fmt.Printf("test1:%p\n", test1)
	fmt.Println("slice4:", slice4[0])

	fmt.Printf("&slice地址append before：%p\n", &slice[0])
	slice = append(slice, slice2...)

	slice[0] = 0
	slice[2] = 9

	arr[0] = 11

	fmt.Printf("&slice地址append after：%p\n", &slice[0])

	fmt.Printf("slice:%v\n", slice)
	fmt.Printf("arr2:%v\n", arr2)
	fmt.Printf("arr:%v\n", arr)

	var test2 Test

	test2.Name = "jiang"
	test2.Age = 31

	c(test2)

	fmt.Println("====================")
	var arr10 [3]int = [3]int{1, 2, 3}
	// var arr11 [3]int = [3]int{1, 2, 3}
	slice10 := arr10[:]
	fmt.Println("====================slice10:", slice10)
	arr10[2] = 33
	fmt.Println("====================arr10:", arr10)
	fmt.Println("====================slice10:", slice10)
	slice10[2] = 333
	fmt.Println("====================arr10:", arr10)
	fmt.Println("====================slice10:", slice10)

	slice10 = append(slice10, 444)
	fmt.Println("====================slice10-append-444")
	fmt.Println("====================slice10:", slice10)
	slice10[2] = 333333
	fmt.Println("====================slice10-重新赋值2-333333:")
	fmt.Println("====================arr10:", arr10)
	fmt.Println("====================slice10:", slice10)

	type A struct {
		name string
	}
	var a A
	a.name = "cccc"
	fmt.Println("aaaaaa:", a.name)

	type D interface {
	}
	var ddd D
	ddd = nil
	fmt.Println(ddd)

	var iii string
	iii = "ccc"
	fmt.Println(iii)

}







func c(aa interface{}) {
	fmt.Println("ccc_interface")
}

//  *
// ***
//*****
func x(c int) {
	m := []map[string]string{
		{"ccc": "ccc"},
		{"bbb": "ccc"},
	}
	fmt.Println(m)

	for i := 1; i <= c; i++ {
		// 在打印之前打印空格
		for k := 1; k <= c-i; k++ {
			fmt.Print(" ")
		}
		// j表示每层打印多少*
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
