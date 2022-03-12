package main

import "fmt"

func main() {
	problem1()
}

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
