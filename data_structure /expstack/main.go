package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int     // 表示我们栈最大可以存放数个数
	Top    int     // 表示栈顶，因为栈顶固定，因此我们直接用top
	arr    [20]int // 数组模拟栈
}

func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

// 变量栈
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	// 判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	// 先取值，再this.Top --
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

// 判断一个字符是不是一个运算符[+-*/]
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 定义运算符优先级
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

// 运算方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

func main() {
	// 数栈
	numStack := &Stack{
		MaxTop: 20, // 表示最多存放20个数到栈中
		Top:    -1, // 当栈顶为-1，表示栈为空
	}
	// 符号栈
	operStack := &Stack{
		MaxTop: 20, // 表示最多存放20个数到栈中
		Top:    -1, // 当栈顶为-1，表示栈为空
	}

	exp := "6-2"

	// 定义一个index ，帮助扫描exp
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""
	for {
		ch := exp[index : index+1]
		temp := int([]byte(ch)[0])  // 就是对应字符串对应的asciI码
		if operStack.IsOper(temp) { // 说明是符号
			// 判断是否为空栈
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				// 如果发现operstack栈顶的运算符大的优先级大于等于当前准备入栈的运算符优先级
				// 就从符号栈pop出，并从数栈中pop出两个数进行运算，运算后的结果再重新入栈到数栈，符号再入符号栈
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					numStack.Push(result)
					operStack.Push(temp)
				}
			}
		} else {
			// 处理多位数思路
			//1 定义一个变量，keepnum string ,做拼接
			keepNum += ch
			//2 每次要向index的前面字符测试一下，看看是不是运算符，然后处理
			// 如果已经到表达式最后，直接将keepnum
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(ch, 10, 64)
				numStack.Push(int(val))
			} else {
				// 向index后面测试看看是不是运算符号
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(ch, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
			val, _ := strconv.ParseInt(ch, 10, 64)
			numStack.Push(int(val))
		}
		if index+1 == len(exp) {
			break
		}
		index++
	}

	// 如果扫描表达式完毕，依次从符号栈中取出符号，然后从数栈取出两个数
	// 运算后的结果，入数栈，直到符号栈为空
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		numStack.Push(result)
	}

	// 如果我们表达式算法没有问题，表达式没有问题，则结果就是numstack最后数
	res, _ := numStack.Pop()
	fmt.Println(res)

}
