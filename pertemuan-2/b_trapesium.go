// package main

// import "fmt"

// func cariTrapesium(base1, base2, height float64) float64 {
// 	return 0.5 * (base1 + base2) * height
// }

// func main() {
// 	base1 := 5.0
// 	base2 := 7.0
// 	height := 8.0

// 	area := cariTrapesium(base1, base2, height)
// fmt.Printf("Luas trapesium dengan alas %.2f dan %.2f tinggi %.2f yaitu %.2f\n", base1, base2, height, area)
// }

package main

import "fmt"

func main() {
	var a, b, t int
	fmt.Println("Masukan nilai A :")
	fmt.Scanln(&a)
	fmt.Println("Masukan nilai B :")
	fmt.Scanln(&b)
	fmt.Println("Masukan nilai T :")
	fmt.Scanln(&t)

	Luas := 0.5 * float64(a+b) * float64(t)
	fmt.Print("Luas trapesium adalah ")
	fmt.Print(Luas)
}
