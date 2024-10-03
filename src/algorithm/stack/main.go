package main

import (
	"errors"
	"fmt"
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

func main() {
	stack := &Stack{
		MaxTop: 5,  //最大存五个
		Top:    -1, //-1表示空栈
	}

	//存
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	//显示
	stack.List()
}
