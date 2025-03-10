package main

import "math/rand"

func getRandomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}
