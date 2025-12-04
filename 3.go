package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func listPrime(n int) {
	fmt.Printf("Cac so nguyen to < %d la : ", n)
	for i := 2; i < n; i++ {
		if isPrime(i) {
			fmt.Print(i, ", ")
		}
	}
	fmt.Println()
}
