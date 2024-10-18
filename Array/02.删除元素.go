package main

func helper(nums []int, target int) int {
	slow := 0
	fast := 0
	for ; fast < len(nums); fast++ {
		if nums[fast] == target {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}
	return slow
}
