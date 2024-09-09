package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子

	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//转稀疏数组
	//遍历chessMap,如果元素的值不为0（创建一个node结构体），将其放入对应的切片
	var sparseArr []ValNode

	//记录元素的原始二维数组的规模，行列默认值
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个ValNode值节点
				valNode = ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//输出稀疏数组
	fmt.Println("当前的稀疏数组是：：：：：：：：")
	for i, valNode := range sparseArr {
		fmt.Printf("%d:%d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	/*-------------*/

	//将稀疏数组恢复原始的数组
	//创建一个原始数组
	var chessMap2 [11][11]int

	for i, valNode := range sparseArr {
		if i != 0 { //跳过第一行记录的值
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}

	//打印
	fmt.Println("恢复后的原始数组：：：：：：")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
