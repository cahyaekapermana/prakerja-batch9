package main

import "fmt"

func cariTrapesium(base1, base2, height float64) float64 {
	return 0.5 * (base1 + base2) * height
}

func main() {
	base1 := 5.0
	base2 := 7.0
	height := 8.0

	area := cariTrapesium(base1, base2, height)
	fmt.Printf("Luas trapesium dengan alas %.2f dan %.2f tinggi %.2f yaitu %.2f\n", base1, base2, height, area)
}
