package main

import "fmt"

/*
有一个公司，当有新员工报道时，要求该员工的信息加入id,性别，年龄，住址，输入该
员工的id时，要求查到该员工的所有信息
*/
type Emp struct {
	Id   int
	Name string
	Next *Emp
}
type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head //辅助指针
	var pre *Emp = nil
	if cur == nil {
		this.Head = emp
		return
	}
	for {
		if cur != nil {
			if cur.Id > emp.Id {
				//找到位置
				break
			}
			pre = cur //保持同步
			cur = cur.Next
		} else {
			break
		}
	}
	//退出时，是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->>", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数确定添加该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	this.LinkArr[linkNo].Insert(emp)
}

// 散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7
}

// 显示
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}
func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for {
		fmt.Println("---雇员菜单---")
		fmt.Println("input   添加雇员")
		fmt.Println("show    显示雇员")
		fmt.Println("find    查找雇员")
		fmt.Println("exit     退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员的name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.ShowAll()
		case "exit":
		default:
			fmt.Println("输入错误")

		}
	}
}
