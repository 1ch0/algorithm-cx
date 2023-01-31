package main

import "fmt"

func main() {
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 || len(prices) > 10e5 {
		return 0
	}
	minI, minV, maxV := 0, prices[0], prices[0]
	for i, v := range prices {
		if minV > v {
			minI = i
			minV = v
			maxV = v
		}
		if maxV < v && i > minI {
			maxV = v
		}
	}
	return maxV - minV
}
