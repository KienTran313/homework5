package main

import "fmt"

func fun4(n int) int {
	if n < 100 || n > 999 {
		return 0
	}
	digti1 := n / 100
	digti2 := (n / 10) % 10
	digti3 := n % 10

	total := (digti1 * digti1 * digti1) + (digti2 * digti2 * digti2) + (digti3 * digti3 * digti3)
	if n == total {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(fun4(1000))
}
