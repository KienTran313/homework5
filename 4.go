package main

import "fmt"

func fun4(n int) int {
	if n <= 100 || n > 1000 {
		fmt.Println("Loi,chi nhap so co ba chu so.")
		return 2
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
