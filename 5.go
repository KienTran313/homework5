package main

func fun5(n int) float64 {
	var sum float64
	for i := 0; i <= n; i++ {
		sum += (2*float64(i) + 1)
	}
	result := 1 / sum
	return result
}

// Test
//
