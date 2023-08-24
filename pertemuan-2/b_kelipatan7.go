package main

import "fmt"

func main() {
	num := 11

	if num%7 == 0 {
		fmt.Printf("%d adalah kelipatan 7.\n", num)
	} else {
		fmt.Printf("%d bukan kelipatan 7.\n", num)
	}
}
