package main

import (
	"bufio"
	"fmt"
	"os"
)

func helper(nums []int) []int {
	res := make([]int, len(nums)+1)
	preSum := 0
	for i, num := range nums {
		preSum += num
		res[i+1] = preSum
	}
	return res
}

func sumRange(nums []int, left int, right int) int {
	preSum := helper(nums)
	return preSum[right+1] - preSum[left]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	vec := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &vec[i])
	}

	var left, right int
	for {
		fmt.Fscan(reader, &left, &right)
		if left == -1 && right == -1 {
			break
		}
		fmt.Fprintln(writer, sumRange(vec, left, right))
	}
}
