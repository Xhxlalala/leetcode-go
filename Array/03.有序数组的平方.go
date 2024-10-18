package main

import "fmt"

func solve(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1
	ans := make([]int, n)
	for i <= j {
		left := nums[i] * nums[i]
		right := nums[j] * nums[j]
		if left <= right {
			ans[k] = right
			j--
		} else {
			ans[k] = left
			i++
		}
		k--
	}
	return ans
}

func main() {
	nums := []int{-5, -1, 3, 4, 5}
	var ans []int
	ans = solve(nums)
	fmt.Println(ans)
	fmt.Printf("%v\n", ans)
}
