// Buat sebuah struct dengan nama Student yang mempunyai properti name dan score dalam bentuk slice kemudian simpan data siswa sebanyak 5 siswa. Setelah 5 siswa dimasukkan maka program menunjukkan skor rata-rata, siswa yang memiliki skor minimum dan maksimal? (implementasikan method)!
package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score float64
}

func main() {
	var students []Student

	// Mengisi data siswa sebanyak 5 siswa
	students = append(students, Student{Name: "Luffy", Score: 90})
	students = append(students, Student{Name: "Zoro", Score: 60})
	students = append(students, Student{Name: "Sanji", Score: 58})
	students = append(students, Student{Name: "Usop", Score: 82})
	students = append(students, Student{Name: "Nami", Score: 78})

	// Menghitung total skor untuk perhitungan rata-rata
	totalScore := 0.0

	// Inisialisasi skor minimum dan maksimal dengan nilai yang ekstrem
	minScore := math.MaxFloat64
	maxScore := -math.MaxFloat64

	// Iterasi untuk menghitung total skor, skor minimum, dan skor maksimal
	for _, student := range students {
		totalScore += student.Score

		if student.Score < minScore {
			minScore = student.Score
		}

		if student.Score > maxScore {
			maxScore = student.Score
		}
	}

	// Menghitung rata-rata
	averageScore := totalScore / float64(len(students))

	// Menampilkan hasil
	fmt.Printf("Skor rata-rata: %.2f\n", averageScore)
	fmt.Printf("Siswa dengan skor minimum: %.2f\n", minScore)
	fmt.Printf("Siswa dengan skor maksimal: %.2f\n", maxScore)
}
