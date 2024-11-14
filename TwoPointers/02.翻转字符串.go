package main

func solution(ss []byte) {
	left := 0
	right := len(ss) - 1
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left++
		right--
	}
}
