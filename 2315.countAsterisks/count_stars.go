package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "l|*e*et|c**o|*de|*"
	fmt.Println(countAsterisks(s))
}

func countAsterisks(s string) int {
	count, isCount := 0, 0

	// 辅助数
	for i := 0; i < len(s); i++ {
		if strings.Contains(string(s[i]), "|") {
			isCount++
		}
		if strings.Contains(string(s[i]), "*") {
			if isCount%2 == 0 {
				count++
			}
		}
	}
	return count
}
