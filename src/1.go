package main 

import "math" 

func getLargestPrime(n int) { 
    var largestPrime = -1 
    for i := 2; i <= n; i++ { 
        if math.IsPrime(i) { 
            largestPrime = i 
        } 
    } 
    return largestPrime 
} 
