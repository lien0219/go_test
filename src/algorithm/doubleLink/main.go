package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //指向前一个节点
	next     *HeroNode //指向下一个节点
}

// 1.双链表插入节点
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
	newHeroNode.pre = temp
}

// 2.链表插入节点，根据no的编号从大到小插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil { //链表最后
			break
		} else if temp.next.no <= newHeroNode.no {
			//newHeroNode插在temp后
			break
		} else if temp.next.no == newHeroNode.no {
			//链表已经有此no就不插入
			flag = false
			break
		}
	}
	if !flag {
		fmt.Println("已经存在no:", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}

// 双向链表删除一个节点
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil { //链表最后
			break
		} else if temp.next.no == id {
			//找到节点
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		//删除
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("要删除的id不存在")
	}
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

// 双链表逆序打印
func ListHeroNode2(head *HeroNode) {
	//辅助节点
	temp := head
	//temp定位到双链表的最后节点
	if temp.next == nil {
		fmt.Println("空链表....")
		return
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	//if temp.next == nil {
	//	fmt.Println("空链表......")
	//	return
	//}
	for {
		fmt.Printf("[%d,%s,%s]==>", temp.no, temp.name, temp.nickname)
		temp = temp.pre
		if temp.pre == nil {
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
	head3 := &HeroNode{
		no:       3,
		name:     "燃灯",
		nickname: "大佛",
	}

	InsertHeroNode(head, head1)
	InsertHeroNode(head, head2)
	InsertHeroNode(head, head3)

	ListHeroNode(head)
	fmt.Println("逆序打印")
	ListHeroNode2(head)
}
