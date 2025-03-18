package main

import "math/rand"

func main() {
	rand.Seed(42) // Random seed to get different results every time
	fmt.Println(rand.Intn(10)) // Random integer between 0 and 9
}
