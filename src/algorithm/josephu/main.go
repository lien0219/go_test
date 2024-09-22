package main

import "fmt"

type Boy struct {
	No   int  //编号
	Next *Boy //下一个孩子的指针
}

func AddBoy(num int) *Boy {

	first := &Boy{} //空节点
	curBoy := &Boy{}

	if num < 1 {
		fmt.Println("num值错误")
		return first
	}
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //构成环形链表
		}
	}

	return first
}

// 显示单向环形链表
func ShowBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}

	curBoy := first
	for {
		fmt.Printf("孩子编号：%d -->>", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

func main() {
	first := AddBoy(5)
	ShowBoy(first)
}
