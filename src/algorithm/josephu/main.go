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
func PlayGame(first *Boy, startNo int, countNum int) {
	//处理空的链表
	if first.Next == nil {
		fmt.Println("链表不能为空！！！")
		return
	}
	//辅助指针，帮助删除
	tail := first
	//让tail指向环形链表的最后一个孩子
	for {
		if tail.Next == first { //最后一个
			break
		}
		tail = tail.Next
	}
	//让first移动到startNo
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	//开始数countNum
	for {
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("孩子编号为：%d 出圈\n", first.No)
		//删除first执行的孩子
		first = first.Next
		tail.Next = first
		//处理圈中只有一个孩子
		if tail == first {
			break
		}
	}
	fmt.Printf("孩子编号为：%d 出圈\n", first.No)
}
func main() {
	first := AddBoy(5)
	ShowBoy(first)
	PlayGame(first, 2, 3)
}
