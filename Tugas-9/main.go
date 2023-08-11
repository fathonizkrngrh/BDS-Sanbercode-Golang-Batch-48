package main

import (
	"fmt"

	"tugas9/geometry"
	"tugas9/phoneinfo"
)

func main() {
	// Soal 1
	fmt.Println("====== Soal 1 ======")

	segitiga := geometry.SegitigaSamaSisi{Alas: 6, Tinggi: 4}
	persegi := geometry.PersegiPanjang{Panjang: 5, Lebar: 3}
	tabung := geometry.Tabung{JariJari: 2.5, Tinggi: 7.0}
	balok := geometry.Balok{Panjang: 4, Lebar: 3, Tinggi: 6}

	var bangunDatar geometry.HitungBangunDatar
	var bangunRuang geometry.HitungBangunRuang

	bangunDatar = segitiga
	fmt.Println("Luas segitiga:", bangunDatar.Luas())
	fmt.Println("Keliling segitiga:", bangunDatar.Keliling())

	bangunDatar = persegi
	fmt.Println("Luas persegi panjang:", bangunDatar.Luas())
	fmt.Println("Keliling persegi panjang:", bangunDatar.Keliling())

	bangunRuang = tabung
	fmt.Println("Volume tabung:", bangunRuang.Volume())
	fmt.Println("Luas permukaan tabung:", bangunRuang.LuasPermukaan())

	bangunRuang = balok
	fmt.Println("Volume balok:", bangunRuang.Volume())
	fmt.Println("Luas permukaan balok:", bangunRuang.LuasPermukaan())

	// Soal 2

	fmt.Println("\n====== Soal 2 ======")

	iphone := phoneinfo.Phone{
		Name:   "iPhone 13",
		Brand:  "Apple",
		Year:   2021,
		Colors: []string{"Black", "White", "Blue", "Red"},
	}

	var info phoneinfo.DisplayInfo
	info = iphone
	fmt.Println(info.GetInfo())

	// Soal 3

	fmt.Println("\n====== Soal 3 ======")
	
	fmt.Println(geometry.LuasPersegi(4, true))
	fmt.Println(geometry.LuasPersegi(8, false))
	fmt.Println(geometry.LuasPersegi(0, true))
	fmt.Println(geometry.LuasPersegi(0, false))

	// Soal 4

	fmt.Println("\n====== Soal 4 ======")

	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	sum := 0

	if angkaPertama, ok := kumpulanAngkaPertama.([]int); ok {
		for _, angka := range angkaPertama {
			sum += angka
		}
	}

	if angkaKedua, ok := kumpulanAngkaKedua.([]int); ok {
		for _, angka := range angkaKedua {
			sum += angka
		}
	}

	output := fmt.Sprintf("%s%d + %d + %d + %d = %d", prefix, kumpulanAngkaPertama.([]int)[0], kumpulanAngkaPertama.([]int)[1], kumpulanAngkaKedua.([]int)[0], kumpulanAngkaKedua.([]int)[1], sum)
	fmt.Println(output)
}
