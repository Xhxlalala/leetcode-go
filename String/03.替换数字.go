package main

import "fmt"

func solution(strByte []byte) string {
	numCount, oldSize := 0, len(strByte)
	for i := 0; i < oldSize; i++ {
		if strByte[i] >= '0' && strByte[i] <= '9' {
			numCount++
		}
	}
	//扩充数组长度
	for i := 0; i < numCount; i++ {
		strByte = append(strByte, []byte("     ")...)
	}
	tmpBytes := []byte("number")
	//双指针从后遍历
	leftP, rightP := oldSize-1, len(strByte)-1
	for leftP < rightP {
		rightShift := 1
		if strByte[leftP] >= '0' && strByte[leftP] <= '9' {
			for i, tmpByte := range tmpBytes {
				strByte[rightP-len(tmpBytes)+i+1] = tmpByte
			}
			rightShift = len(tmpBytes)
		} else {
			strByte[rightP] = strByte[leftP]
		}
		// 移动指针
		leftP--
		rightP -= rightShift
	}
	return string(strByte)
}

func main() {
	var strByte []byte
	fmt.Scanln(&strByte)
	fmt.Println(solution(strByte))
}
