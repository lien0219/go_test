package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode //指向下一个节点
}

// 链表插入节点
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil { //找到最后
			break
		}
		//不断的指向下一个节点
		temp = temp.next
	}
	//新节点加入链表最后
	temp.next = newHeroNode

}

// 显示链表的所有信息
func ListHeroNode(head *HeroNode) {
	//辅助节点
	temp := head
	if temp.next == nil {
		fmt.Println("空链表......")
		return
	}
	for {
		fmt.Printf("[%d,%s,%s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}
func main() {

	//头节点
	head := &HeroNode{}

	//新的HeroNode
	head1 := &HeroNode{
		no:       1,
		name:     "孙悟空",
		nickname: "行者",
	}
	head2 := &HeroNode{
		no:       2,
		name:     "如来",
		nickname: "佛祖",
	}

	//加入
	InsertHeroNode(head, head1)
	InsertHeroNode(head, head2)
	//显示
	ListHeroNode(head)
}
