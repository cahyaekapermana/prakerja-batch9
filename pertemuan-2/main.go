// package main = yang pertama kali dijalankan
package main

import "fmt"

// membuat fungsi main
// run file bisa dengan build .exe dengan "go build main.go":
// atau ketik "go run main.go"
// atau "./main" untuk ditampilkan ke terminal
func main() {

	// age := 11

	// if age == 20 {
	// 	fmt.Println("Muda")
	// } else if age < 20 {
	// 	fmt.Println("Teenager")
	// } else {
	// 	fmt.Println("Tua")
	// }
	// Loop
	// Segitiga siku2
	// for j := 1; j <= 5; j++ {
	// 	// isi
	// 	for i := 1; i <= j; i++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	for j := 1; j <= 5; j++ {
		// isi
		if j == 3 {
			break
		}
		fmt.Println(j)
	}

}
