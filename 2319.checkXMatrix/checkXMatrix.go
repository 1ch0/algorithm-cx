package main

import "fmt"

func main() {
	//grid := [][]int{[[6,0,0,2],[0,5,15,0],[9,5,4,0],[8,0,0,20]]}
	grid := [][]int{{6, 0, 0, 2}, {0, 5, 15, 0}, {9, 5, 4, 0}, {8, 0, 0, 20}}
	//grid := [][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}}
	fmt.Println(checkXMatrix(grid))
}

func checkXMatrix(grid [][]int) bool {
	// 确认为 0 的数据下标值,然后比较
	for i, v := range grid {
		for j, v2 := range v {
			if i == j || len(v)-i-1 == j {
				if v2 == 0 {
					fmt.Println("test")
					fmt.Println(i, v, j, v2)
					return false
				}
				continue
			}
			if v2 != 0 {
				fmt.Println(i, v, j, v2)
				return false
			}
		}
	}
	return true
}
