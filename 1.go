package main

func fun1(n int) int {
	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}
	return total
}
