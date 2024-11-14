package main

func solution(str []byte) string {
	numCount := 0
	oldSize := len(str)
	for i := 0; i < oldSize; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			numCount++
		}
	}

	for i := 0; i < numCount; i++ {
		//每个数字添加5个空格
		str = append(str, []byte("     ")...)
	}
	tmpBytes := []byte("number")

	left, right := oldSize-1, len(str)-1
	for left < right {
		rightShift := 1
		if str[left] >= '0' && str[left] <= '9' {
			for i, tmpByte := range tmpBytes {
				str[right-len(tmpBytes)+i+1] = tmpByte
			}
			rightShift = len(tmpBytes)
		} else {
			str[right] = str[left]
		}
		left--
		right -= rightShift
	}

	return string(str)
}
