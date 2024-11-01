package main

func solve(A []int, B []int, C []int, D []int) int {
	m := make(map[int]int)
	count := 0
	for _, a := range A {
		for _, b := range B {
			m[a+b]++
		}
	}

	for _, c := range C {
		for _, d := range D {
			sum := -c - d
			if val, ok := m[sum]; ok {
				count += val
			}
		}
	}
	return count
}
