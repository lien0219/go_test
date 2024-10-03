package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 排序函数
func SelectSort(arr *[80000]int) {
	//(*arr)[1] = 600
	//将最大值和arr[0]交换
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		//fmt.Printf("第%d次：%v \n", j+1, *arr)
	}
}
func main() {
	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}
	//arr := [5]int{10, 34, 19, 100, 80}
	start := time.Now().Unix()
	SelectSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时：%d \n", end-start)
	fmt.Println("main函数")
	//fmt.Println(arr)
}
