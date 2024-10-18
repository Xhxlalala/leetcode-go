package main

func solve(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	startx, starty := 0, 0
	loop := n / 2
	mid := n / 2
	count := 1
	offset := 1
	for loop > 0 {
		i, j := startx, starty
		// 从左到右
		for j = starty; j < n-offset; j++ {
			ans[i][j] = count
			count++
		}
		// 从上到下
		for i = startx; i < n-offset; i++ {
			ans[i][j] = count
			count++
		}
		// 从右到左
		for ; j > starty; j-- {
			ans[i][j] = count
			count++
		}
		// 从下到上
		for ; i > startx; i-- {
			ans[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}
	if n%2 == 1 {
		ans[mid][mid] = count
	}
	return ans
}
