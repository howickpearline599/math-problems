package main

import "fmt"

func solve(n int) int {
    if n == 0 {
        return 1
    }
    return n * solve(n-1)
}

func main() {
	fmt.Println(solve(5)) // should print 120
}
