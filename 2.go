package main

func fun2(n, m int) int {
	result1 := 1
	for i := 1; i <= m; i++ {
		if n < 0 && m%2 == 0 {
			result1 *= -n
		} else {
			result1 *= n
		}
	}

	result2 := 1
	for i := 1; i <= n; i++ {
		if m < 0 && n%2 == 0 {
			result2 *= -m
		} else {
			result2 *= m
		}
	}
	total := result1 + result2
	return total
}
