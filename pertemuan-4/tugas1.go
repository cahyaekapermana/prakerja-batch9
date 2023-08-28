package main

import (
	"fmt"
)

type Car struct {
	Type    string
	FuelInL float64
}

func (c Car) PerkiraanJarak() float64 {
	// Konversi bahan bakar dari liter ke milliliter
	fuelInMl := c.FuelInL * 4000.0

	// Menghitung perkiraan jarak berdasarkan konsumsi bahan bakar
	// (1.5 L/mill)
	jarak := fuelInMl / 1.5

	return jarak
}

func main() {
	myCar := Car{
		Type:    "Alphard",
		FuelInL: 40.0,
	}

	jarak := myCar.PerkiraanJarak()
	fmt.Printf("Jarak perkiraan yang bisa ditempuh dengan mobil tipe %s dan bahan bakar %.2f L: %.2f mill\n", myCar.Type, myCar.FuelInL, jarak)
}
