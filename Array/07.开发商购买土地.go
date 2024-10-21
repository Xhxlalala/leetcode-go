package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sum := 0
	vec := make([][]int, n)
	for i := range vec {
		vec[i] = make([]int, m)
		for j := range vec[i] {
			fmt.Scan(&vec[i][j])
			sum += vec[i][j]
		}
	}

	result := math.MaxInt32
	count := 0 // 统计遍历过的行

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count += vec[i][j]
			// 遍历到行末尾时候开始统计
			if j == m-1 {
				result = min(result, abs(sum-count-count))
			}
		}
	}

	count = 0 // 统计遍历过的列
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			count += vec[i][j]
			// 遍历到列末尾的时候开始统计
			if i == n-1 {
				result = min(result, abs(sum-count-count))
			}
		}
	}

	fmt.Println(result)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
