package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) Print() {
	fmt.Printf("Name:%v，Age:%d；", s.Name, s.Age)
}
func test01(arg interface{}) {
	stu := arg.(Student)
	stu.Print()
}

func main() {
	var stu Student
	stu = Student{
		Name: "蒋生平",
		Age:  31,
	}
	test01(stu)
}
