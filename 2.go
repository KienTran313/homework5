package main

func fun2(n, m int) int {
	result1 := 1
	for i := 1; i <= m; i++ {
		// if n < 0 && m%2 == 0 {
		// 	result1 *= -n
		// } else {
		result1 *= n
		// }
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

func func3(n int, m int) float32 {
	var nm float32 = 1.0

	// Tinh n ^ m
	// - m > 0; nm = n * n * n* .... *n
	// - m < 0; nm = 1/n * 1/n * 1*n = 1/ n^-m
	//

	m1 := float32(m)
	n1 := float32(n)
	if m < 0 {
		m1 = -float32(m)
		n1 = 1 / float32(n)
	}
	var i float32 = 0
	for ; i < m1; i++ {
		nm *= n1
	}

	// Tinh m ^ n
	// m * m * ... *m
	// 1/m * 1/m ... * 1/m
	//

	n1 = float32(n)
	m1 = float32(m)
	if n < 0 {
		n1 = -float32(n)
		m1 = 1 / float32(m)
	}

	var mn float32 = 1
	i = 0
	for ; i < n1; i++ {
		mn *= m1
	}

	return nm + mn
}

/*
* n = -3, m = 4
*
* result1 = 1
*
* i = 1; result1 *= -n = 1 * - (-3) = 3
* i = 2; result *= ....... = 3 * 3 = 9
* i = 3; result = 9 * 3 = 27
* i = 4; result = 27 * 3 = 81
* i = 5
* result1 = 81
*
* result2 = 1
* i = 1; XXXXX
*
* total = result 1 + 2 = 81 + 1 = 82
*
*
*
* n = -3, m = 4
*
* result1 = 1
*
* i = 1; result1 *= -3 = -3
* i = 2; result1 = -3 * -3 = 9
* i = 3; result1 = 9 * -3 = -27
* i = 4l result1 = -27 * -3 = 81
* */
