package _009_isPalindrome

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	// 负数默认不是回文数
	if x < 0 {
		return false
	}
	// 将数据拆分为字符数组，比较数据元素
	v := strconv.Itoa(x)
	v1 := strings.Split(v, "")
	if len(v1)%2 != 0 {
		// 奇数
		for i, v := range v1[:(len(v1)-1)/2] {
			if v != v1[len(v1)-i-1] {
				return false
			}
		}
	} else {
		// 偶数
		for i, v := range v1[:len(v1)/2] {
			if v != v1[len(v1)-i-1] {
				return false
			}
		}
	}
	// 十进制左移，右移
	return true
}
