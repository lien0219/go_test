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

// 删除
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	//空链表
	if temp.next == nil {
		fmt.Println("空环形链表，不能删除")
		return head
	}
	//如果只有一个节点
	if temp.next == head {
		temp.next = nil
		return head
	}
	//将helper定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//包含两个以上的节点
	flag := true
	for {
		if temp.next == head { //最后一个
			break
		}
		if temp.no == id { //找到
			if temp == head {
				head = head.next
			}
			helper.next = temp.next
			fmt.Printf("猫的id:%d\n", id)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫的id:%d\n", id)
		} else {
			fmt.Printf("没有对应的id:%d \n", id)
		}
	}

	return head
}
func main() {
	//环形链表的头节点
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "梨花",
	}
	cat2 := &CatNode{
		no:   2,
		name: "橘猫",
	}
	//cat3 := &CatNode{
	//	no:   3,
	//	name: "布偶",
	//}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	//InsertCatNode(head, cat3)
	ListCircleLink(head)
	head = DelCatNode(head, 1)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	ListCircleLink(head)
}
