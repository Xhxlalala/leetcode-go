package main

func reverse(s []byte, l, r int) {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func solution(str string, k int) string {
	ss := []byte(str)
	length := len(ss)
	reverse(ss, 0, length-1)
	reverse(ss, 0, k-1)
	reverse(ss, k, length-1)
	return string(ss)
}
