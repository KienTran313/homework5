package main

func fun5(n int) float64 {
	var sum float64
	for i := 0; i <= n; i++ {
		sum += 1 / (2*float64(i) + 1)
	}
	return sum
}

// = 1/ (1 + 3 + .. + 2n + 1)
