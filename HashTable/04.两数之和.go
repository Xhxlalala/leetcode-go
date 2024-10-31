package main

func solve(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, val := range nums {
		// map:{key:val--数组元素, value:idx--数组元素对应下标}
		if preIdx, ok := m[target-val]; ok {
			return []int{preIdx, idx}
		} else {
			m[val] = idx
		}
	}
	return []int{}
}
