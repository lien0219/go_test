package main

import (
	"errors"
	"fmt"
	"os"
)

// 结构体管理队列
type Queue struct {
	maxSize int
	array   [4]int //模拟队列
	front   int    // 表示指向队列首
	real    int    //表示队列尾
}

// 添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	//队列是否已满
	if this.real == this.maxSize-1 {
		return errors.New("queue fill")
	}

	this.real++ //real后移
	this.array[this.real] = val

	return
}

// 显示队列，找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	for i := this.front + 1; i <= this.real; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

// 取数据
func (this *Queue) GetQueue() (val int, err error) {
	if this.real == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		real:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1.输入add  表示添加数据到队列")
		fmt.Println("2.输入get  表示从队列获取数据")
		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出一个数：", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
