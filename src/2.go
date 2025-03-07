package main

import "fmt"

func solveProblem(n int) int {
	if n < 2 {
		return n
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return -1 // not prime
		}
	}
	return n // prime
}
