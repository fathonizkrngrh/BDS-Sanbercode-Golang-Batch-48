package main

import (
	"fmt"
	"math"
)

// Soal 1

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2*(p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2*math.Pi*t.jariJari*t.tinggi + 2*math.Pi*t.jariJari*t.jariJari
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2*float64(b.panjang*b.lebar+b.panjang*b.tinggi+b.lebar*b.tinggi)
}

// Soal 2

type phone struct {
	name   string
	brand  string
	year   int
	colors []string
}

type displayInfo interface {
	getInfo() string
}

func (p phone) getInfo() string {
	info := fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolor : %v", p.name, p.brand, p.year, p.colors)
	return info
}

// Soal 3

func luasPersegi(sisi int, showDetails bool) interface{} {
	if sisi == 0 {
		if showDetails {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi

	if showDetails {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}

	return luas
}

func main() {
	// Soal 1

	fmt.Println("====== Soal 1 ======")

	segitiga := segitigaSamaSisi{alas: 6, tinggi: 4}
	persegi := persegiPanjang{panjang: 5, lebar: 3}
	tabung := tabung{jariJari: 2.5, tinggi: 7.0}
	balok := balok{panjang: 4, lebar: 3, tinggi: 6}

	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang

	bangunDatar = segitiga
	fmt.Println("Luas segitiga:", bangunDatar.luas())
	fmt.Println("Keliling segitiga:", bangunDatar.keliling())

	bangunDatar = persegi
	fmt.Println("Luas persegi panjang:", bangunDatar.luas())
	fmt.Println("Keliling persegi panjang:", bangunDatar.keliling())

	bangunRuang = tabung
	fmt.Println("Volume tabung:", bangunRuang.volume())
	fmt.Println("Luas permukaan tabung:", bangunRuang.luasPermukaan())

	bangunRuang = balok
	fmt.Println("Volume balok:", bangunRuang.volume())
	fmt.Println("Luas permukaan balok:", bangunRuang.luasPermukaan())

	// Soal 2

	fmt.Println("\n====== Soal 2 ======")

	iphone := phone{
		name:   "iPhone 13",
		brand:  "Apple",
		year:   2021,
		colors: []string{"Black", "White", "Blue", "Red"},
	}

	var info displayInfo
	info = iphone
	fmt.Println(info.getInfo())

	// Soal 3

	fmt.Println("\n====== Soal 3 ======")
	
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

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
