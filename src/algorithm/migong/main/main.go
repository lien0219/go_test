package main

import "fmt"

// 地图：myMap *[8][7]int
// 地图上的点：i int, j int
func SetWay(myMap *[8][7]int, i int, j int) bool {
	if myMap[6][5] == 2 {
		return true
	} else {
		//继续找
		if myMap[i][j] == 0 { //可以走
			//假设可以走
			//上下左右
			//下右上左
			myMap[i][j] = 2
			if SetWay(myMap, i-1, j) { //上
				return true
			} else if SetWay(myMap, i+1, j) { //下
				return true
			} else if SetWay(myMap, i, j-1) { //左
				return true
			} else if SetWay(myMap, i, j+1) { //右
				return true
			} else { //死路
				myMap[i][j] = 3
				return false
			}
		} else { //不能走,1是墙
			return false
		}
	}
}
func main() {
	//二维数组模拟迷宫
	var myMap [8][7]int

	//地图的最上和最下设置1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	//最左和最右设置1
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	SetWay(&myMap, 1, 1)
	fmt.Println("探测完成：")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], "")
		}
		fmt.Println()
	}
}
