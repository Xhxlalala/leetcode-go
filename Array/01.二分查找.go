package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)>>2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	res := search(nums, 4)
	fmt.Println(res)
	res = search2(nums, 4)
	fmt.Println(res)
}
