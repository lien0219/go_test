package main

import (
	"fmt"
	"os"
)

/*
有一个公司，当有新员工报道时，要求该员工的信息加入id,性别，年龄，住址，输入该
员工的id时，要求查到该员工的所有信息
*/
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (this *Emp) ShowMe() {
	fmt.Printf("链表%d找到该雇员%d\n", this.Id%7, this.Id)
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

// 查找
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}
func (this *HashTable) FindById(id int) *Emp {
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
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
		case "find":
			fmt.Println("请输入id号：")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Printf("id-%d 的雇员不存在\n", id)
			} else {
				//显示雇员信息
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")

		}
	}
}
