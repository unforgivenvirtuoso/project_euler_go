package main

import (
	"fmt"
	"strconv"
)

func main() {
	problem6()
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

// Find the largest palindrome made from the product of two 3-digit numbers.
func problem4() {
	var largestPalindrome int

	for first := 999; first > 99; first-- {
		for second := 999; second > 99; second-- {
			value := strconv.Itoa(first * second)

			isPalindrome := false
			for i := 0; i < len(value)/2; i++ {
				if value[i] == value[len(value)-i-1] {
					isPalindrome = true
				} else {
					isPalindrome = false
					break
				}
			}

			if isPalindrome {
				palindrome, _ := strconv.Atoi(value)
				if palindrome > largestPalindrome {
					largestPalindrome = palindrome
				}
			}

		}
	}
	fmt.Println(largestPalindrome)
}

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
func problem5() {
	n := 21
	answer_found := false
	divisible := false

	for !answer_found {
		for i := 1; i < 21; i++ {
			if n%i == 0 {
				divisible = true
			} else {
				n++
				i = 1
			}
		}
		if divisible {
			answer_found = true
		}

	}
	fmt.Println(n)
}

func problem6() {
	sum_of_squares := 0
	squares_of_sum := 0
	dif := 0

	for i := 0; i < 101; i++ {
		sum_of_squares += (i * i)

		squares_of_sum += i
	}
	squares_of_sum = squares_of_sum * squares_of_sum

	dif = squares_of_sum - sum_of_squares

	fmt.Println(dif)

}
