package main

func reverse(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func solution(words string) string {
	// 去除多余空格
	ss := []byte(words)
	slow, fast := 0, 0
	length := len(ss)
	for length > 0 && fast < length && ss[fast] == ' ' {
		// 去除头部空格
		fast++
	}
	for ; fast < length; fast++ {
		if fast-1 > 0 && ss[fast-1] == ss[fast] && ss[fast] == '' {
			// 去除单词间多余空格
			continue
		}
		ss[slow] = ss[fast]
		slow++
	}
	if slow-1 > 0 && ss[slow-1] == ' ' {
		// 去除尾部空格
		ss = ss[:slow-1]
	} else {
		ss = ss[:slow]
	}
	// 翻转整个字符串
	reverse(ss)
	// 翻转每个单词
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
