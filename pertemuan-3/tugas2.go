// buatlah sebuah program yang dapat menghitung berapa banyak sebuah string yang sama didalam sebuah slice!
package main

import (
	"fmt"
)

func countStringOccurrences(slice []string, target string) int {
	count := 0
	for _, value := range slice {
		if value == target {
			count++
		}
	}
	return count
}

func main() {
	stringSlice := []string{"Luffy", "Zoro", "Sanji", "Usop", "Shanks", "Shanks"}
	targetString := "Shanks"

	occurrences := countStringOccurrences(stringSlice, targetString)
	fmt.Printf("Jumlah string '%s' dalam slice: %d\n", targetString, occurrences)
}
