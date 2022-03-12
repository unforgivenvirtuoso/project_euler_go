package main

import "fmt"

func main() {
	problem3()
}

// Find the sum of all the multiples of 3 or 5 below 1000.
func problem1() {

	n := 1000
	sum := 0

	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

}

// By considering the terms in the Fibonacci sequence whose values do not exceed four million,
// find the sum of the even-valued terms.
func problem2() {
	current := 1
	next := 2
	even := 0

	for current < 4000000 {
		fib := current + next

		if fib%2 == 0 {
			even += fib
		}

		current = next
		next = fib
	}
	fmt.Println(even)
}

// What is the largest prime factor of the number 600851475143
func problem3() {
	n := 600851475143
	prime := 2

	for prime < n {
		for n%prime == 0 {
			n = n / prime
		}
		prime++
	}
	fmt.Println(prime)

}
