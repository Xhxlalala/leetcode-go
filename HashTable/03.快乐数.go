package main

func solve(n int) bool {
	m := make(map[int]bool)
	//只关心值，所以采用单值取map值
	for n != 1 && !m[n] {
		m[n] = true
		n = getSum(n)
	}
	return n == 1
}

func getSum(n int) int {
	var sum int
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
