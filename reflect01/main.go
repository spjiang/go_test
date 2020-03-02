package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}
type Teacher struct {
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

func test02(a interface{}) {
	ra := reflect.ValueOf(a)
	ri := ra.Interface()

	switch a.(type) {
	case Teacher:
		rt := ri.(Teacher)
		fmt.Println("这是老师：", rt.Name)
	case Student:
		rs := ri.(Student)
		fmt.Println("这是学生：", rs.Name)
	}
}

func test03(a interface{}) {
	switch a.(type) {
	case Teacher:
		rt := a.(Teacher)
		fmt.Println("这是老师：", rt.Name)
	case Student:
		rs := a.(Student)
		fmt.Println("这是学生：", rs.Name)
	}
}

func main() {
	var s Student
	s = Student{
		Name: "蒋学生",
		Age:  31,
	}
	var t Teacher
	t = Teacher{
		Name: "蒋老师",
		Age:  31,
	}
	test03(s)
	test03(t)
}
