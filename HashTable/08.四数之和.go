package main

import "sort"

func solution(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			b := nums[j]
			if j > i+1 && b == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				c := nums[l]
				d := nums[r]
				sum := a + b + c + d
				if sum == target {
					res = append(res, []int{a, b, c, d})
					for l < r && nums[l+1] == c {
						l++
					}
					for l < r && nums[r-1] == d {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
