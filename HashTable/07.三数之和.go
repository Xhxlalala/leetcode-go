package main

import "sort"

func solution(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if a > 0 {
			break
		}
		if i > 0 && a == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			b, c := nums[l], nums[r]
			if a+b+c == 0 {
				res = append(res, []int{a, b, c})
				for l < r && nums[l] == b {
					l++
				}
				for l < r && nums[r] == c {
					r--
				}
			} else if a+b+c < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
