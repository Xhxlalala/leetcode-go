package main

func solve(nums1, nums2 []int) []int {
	set := make(map[int]struct{}, 0)
	for _, v := range nums1 {
		//go语言中从 map 获取值会返回两个结果：值本身和一个布尔值（表示键是否存在）
		//如果键不存在，设置为set[v] = struct{}{}, 表示键存在(键是关键，不关心值)
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	res := make([]int, 0)
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
