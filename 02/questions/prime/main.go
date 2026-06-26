package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	primes := findPrimes(n)
	fmt.Printf("Found %d primes up to %d\n", len(primes), n)
	fmt.Println("Primes:", primes)
}

func findPrimes(n int) []int {
	// YOUR CONCURRENT IMPLEMENTATION HERE
	// Choose your strategy and implement it
	return []int{}
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}