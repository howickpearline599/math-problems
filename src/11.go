package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 10
	randomNumber := rand.Intn(10) + 1

	fmt.Println("The random number is:", randomNumber)
}
