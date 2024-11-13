package main

func solution(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	next := make([]int, n)

	j := -1
	next[0] = j
	//j+1 表示当前已匹配的前缀长度
	//当不匹配时，j = next[j] 表示回退到更短的可能匹配位置
	//当匹配时，j++ 表示匹配长度加1
	for i := 1; i < n; i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
	//如果字符串可以由重复单位组成，那么它一定存在相同的前后缀
	//next[n-1]+1表示最长相等前后缀的长度
	//n-(next[n-1]+1)表示重复子串的长度
	if next[n-1] != -1 && n%(n-(next[n-1]+1)) == 0 {
		return true
	}
	return false
}
