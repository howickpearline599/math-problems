  package main

import (
"fmt"
"math/rand"
)

func main() {
// Generate a random number between 1 and 10
randomNumber := rand.Intn(10) + 1

// Print the random number to the console
fmt.Println("The random number is", randomNumber)
}