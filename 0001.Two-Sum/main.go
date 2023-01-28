package main

import "fmt"

func main() {
	t := []int{1, 2}
	s := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(s, t)

}

func twoSum(nums []int, target int) []int {
	var answer []int

	if len(nums) < 2 || len(nums) > 10e4 {
		return nil
	}
	for k, v := range nums {
		for k2, v2 := range nums[k+1:] {
			if (v + v2) == target {
				answer = append(answer, k, k2+k+1)
				return answer
			}
		}
	}
	return nil
}
