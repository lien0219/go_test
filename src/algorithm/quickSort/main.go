package main

import "fmt"

func QuickSort(left int, right int, array *[6]int) {
	l := left
	r := right
	//pivot中轴
	pivot := array[(left+right)/2]
	temp := 0
	for l < r {
		for array[l] < pivot {
			l++
		}
		for array[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	//左递归
	if left < r {
		QuickSort(left, r, array)
	}
	//右递归
	if right > l {
		QuickSort(l, right, array)
	}
}
func main() {
	arr := [6]int{-9, 78, 0, 23, -567, 70}
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println("main....")
	fmt.Println(arr)
}
