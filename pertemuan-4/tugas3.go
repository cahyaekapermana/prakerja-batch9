// Buatlah program di Golang untuk menemukan nilai maksimum serta minimum di antara 6 angka inputan. Gunakan multiple return fungsi, pointer untuk referencing maupun deferencing!
package main

import (
	"fmt"
)

func findMinMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	var inputNumbers [6]int

	fmt.Println("Masukkan 6 angka:")
	for i := 0; i < 6; i++ {
		fmt.Printf("Angka ke-%d: ", i+1)
		fmt.Scan(&inputNumbers[i])
	}

	minValue, maxValue := findMinMax(inputNumbers[:])

	fmt.Printf("Nilai minimum: %d\n", minValue)
	fmt.Printf("Nilai maksimum: %d\n", maxValue)

	// Menggunakan pointer untuk mengganti nilai minimum dan maksimum
	var ptrMin *int = &minValue
	var ptrMax *int = &maxValue

	*ptrMin = 99
	*ptrMax = 101

	fmt.Printf("Nilai minimum setelah dereferensi: %d\n", minValue)
	fmt.Printf("Nilai maksimum setelah dereferensi: %d\n", maxValue)
}
