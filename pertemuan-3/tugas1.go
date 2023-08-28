// Buatlah sebuah program menggabungkan 2 array yang diberikan, dan jangan sampai terdapat nama yang sama di data yang sudah tergabung tadi!
package main

import (
	"fmt"
)

func mergeArrays(arr1 []string, arr2 []string) []string {
	merged := make([]string, 0, len(arr1)+len(arr2))
	seenNames := make(map[string]bool)

	for _, name := range arr1 {
		if !seenNames[name] {
			seenNames[name] = true
			merged = append(merged, name)
		}
	}

	for _, name := range arr2 {
		if !seenNames[name] {
			seenNames[name] = true
			merged = append(merged, name)
		}
	}

	return merged
}

func main() {
	array1 := []string{"Luffy", "Zoro", "Sanji"}
	array2 := []string{"Zoro", "Nami", "Robin"}

	mergedArray := mergeArrays(array1, array2)
	fmt.Println("Merged Array:", mergedArray)
}
