package main

func solution(words string) string {
	ss := []byte(words)
	slow, fast := 0, 0
	length := len(ss)
	for length > 0 && fast < length && ss[fast] == ' ' {
		fast++
	}
	for ; fast < length; fast++ {
		if fast-1 > 0 && ss[fast-1] == ss[fast] && ss[fast] == ' ' {
			continue
		}
		ss[slow] = ss[fast]
		slow++
	}
	if slow-1 > 0 && ss[slow-1] == ' ' {
		ss = ss[:slow-1]
	} else {
		ss = ss[:slow]
	}
	reverse(ss)
	wordStart := 0
	for wordStart < len(ss) {
		wordEnd := wordStart
		for wordEnd < len(ss) && ss[wordEnd] != ' ' {
			wordEnd++
		}
		reverse(ss[wordStart:wordEnd])
		wordStart = wordEnd + 1
	}
	return string(ss)
}

func reverse(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
