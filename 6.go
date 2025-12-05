package main

import "fmt"

func fun6(n int) float64 {
	var result float64
	for i := 1; i <= n; i++ {
		sum := 0.0
		for j := 1; j <= i; j++ {
			sum += float64(j)
			fmt.Println(sum, ", ")
		}
		result += 1 / sum
	}
	return result
}

func main() {
	fmt.Print(fun6(5))
}

// n = 4
// i = 1 j = 1
// i = 2 j =1;j=2
// i = 3 j=1;j=2;j=3
// i = 4 j =1;j=2;j=3;j=4
//
