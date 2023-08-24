package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	num := 7
	if isPrime(num) {
		fmt.Printf("%d adalah bilangan prima.\n", num)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", num)
	}
}
