package main

import "fmt"

type CatNode struct {
	no   int //编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加的第一个
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //构成环形
		fmt.Println(newCatNode, "加入到环形链表")
		return
	}
	//临时变量找最后节点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到环形链表
	temp.next = newCatNode
	newCatNode.next = head
}

// 输出环形链表
func ListCircleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("环形链表为空！！！")
		return
	}
	for {
		fmt.Printf("猫的信息为：[id=%d name=%s] ->>\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}
func main() {
	//环形链表的头节点
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "梨花",
	}

	InsertCatNode(head, cat1)
	ListCircleLink(head)
}
