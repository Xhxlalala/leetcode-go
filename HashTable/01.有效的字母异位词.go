package main

// 使用数组来做哈希的题目，是因为题目限制了数值的大小
func solve(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := [26]int{}
	for _, v := range s {
		record[v-'a']++
	}
	for _, v := range t {
		record[v-'a']--
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}
