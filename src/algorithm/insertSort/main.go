package main

import "fmt"

func InsertSort(arr *[5]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d插入结果：%v \n", i, *arr)
	}
}
func main() {
	arr := [5]int{23, 0, 12, 56, 34}
	fmt.Println("原始数组：", arr)
	InsertSort(&arr)
	fmt.Println("main函数")
	fmt.Println(arr)
}
