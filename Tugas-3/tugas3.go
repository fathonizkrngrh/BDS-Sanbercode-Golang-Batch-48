package main

import (
	"fmt"
	"strconv"
)

func main() {
	//  Soal 1

	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	luasPersegiPanjang := panjang * lebar
	kelilingPersegiPanjang := 2 * (panjang + lebar)

	luasSegitiga := (alas * tinggi) / 2

	// Tampilkan hasil perhitungan
	fmt.Printf("Luas Persegi Panjang: %d\n", luasPersegiPanjang)
	fmt.Printf("Keliling Persegi Panjang: %d\n", kelilingPersegiPanjang)
	fmt.Printf("Luas Segitiga: %d\n", luasSegitiga)
}