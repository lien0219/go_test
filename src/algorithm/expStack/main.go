package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 模拟栈的使用
type Stack struct {
	MaxTop int    //最大可以存放数
	Top    int    //栈顶
	arr    [5]int //栈
}

func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	//存入
	this.arr[this.Top] = val
	return
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

// 遍历栈
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的遍历：")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

// 判断字符
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
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

// 运算符优先级
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}
func main() {
	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	exp := "3+2*6-1"
	//index,辅助扫描exp
	index := 0
	//配合运算符
	num1 := 0
	num2 := 0
	oper := 0
	result := 0

	keepNum := ""

	for {
		ch := exp[index : index+1]
		temp := int([]byte(ch)[0]) //字符对应的ASCil码
		if operStack.IsOper(temp) {
			if operStack.Top == -1 { //空栈
				operStack.Push(temp)
			} else {
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = numStack.Pop()
					result = numStack.Cal(num1, num2, oper)
					//结果重新入栈
					numStack.Push(result)
					//符号入栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else {
			keepNum += ch
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
			//val, _ := strconv.ParseInt(ch, 10, 64)
			//numStack.Push(int(val))
		}
		//扫描
		if index+1 == len(exp) {
			break
		}
		index++
	}
	//入数栈，直到符号栈为空
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = numStack.Pop()
		result = numStack.Cal(num1, num2, oper)
		//结果重新入栈
		numStack.Push(result)
	}
	res, _ := numStack.Pop()
	fmt.Printf("表达式为：%s=%v", exp, res)
}
