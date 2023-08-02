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

	fmt.Println("====== Jawaban 1 ======")
	fmt.Printf("Luas Persegi Panjang: %d\n", luasPersegiPanjang)
	fmt.Printf("Keliling Persegi Panjang: %d\n", kelilingPersegiPanjang)
	fmt.Printf("Luas Segitiga: %d\n", luasSegitiga)

	// Soal 2

	var nilaiJohn = 80
    var nilaiDoe = 50

	var indeksJohn string
    if nilaiJohn >= 80 {
        indeksJohn = "A"
    } else if nilaiJohn >= 70 && nilaiJohn < 80 {
        indeksJohn = "B"
    } else if nilaiJohn >= 60 && nilaiJohn < 70 {
        indeksJohn = "C"
    } else if nilaiJohn >= 50 && nilaiJohn < 60 {
        indeksJohn = "D"
    } else {
        indeksJohn = "E"
    }

    var indeksDoe string
    if nilaiDoe >= 80 {
        indeksDoe = "A"
    } else if nilaiDoe >= 70 && nilaiDoe < 80 {
        indeksDoe = "B"
    } else if nilaiDoe >= 60 && nilaiDoe < 70 {
        indeksDoe = "C"
    } else if nilaiDoe >= 50 && nilaiDoe < 60 {
        indeksDoe = "D"
    } else {
        indeksDoe = "E"
    }

	fmt.Println("\n====== Jawaban 2 ======")
	fmt.Println("Indeks nilai John:", indeksJohn)
    fmt.Println("Indeks nilai Doe:", indeksDoe)

	// Soal 3

	var tanggal = 30
    var bulan = 10
    var tahun = 2002

	fmt.Println("\n====== Jawaban 3 ======")
    switch bulan {
    case 1:
        fmt.Printf("%d Januari %d\n", tanggal, tahun)
    case 2:
        fmt.Printf("%d Februari %d\n", tanggal, tahun)
    case 3:
        fmt.Printf("%d Maret %d\n", tanggal, tahun)
    case 4:
        fmt.Printf("%d April %d\n", tanggal, tahun)
    case 5:
        fmt.Printf("%d Mei %d\n", tanggal, tahun)
    case 6:
        fmt.Printf("%d Juni %d\n", tanggal, tahun)
    case 7:
        fmt.Printf("%d Juli %d\n", tanggal, tahun)
    case 8:
        fmt.Printf("%d Agustus %d\n", tanggal, tahun)
    case 9:
        fmt.Printf("%d September %d\n", tanggal, tahun)
    case 10:
        fmt.Printf("%d Oktober %d\n", tanggal, tahun)
    case 11:
        fmt.Printf("%d November %d\n", tanggal, tahun)
    case 12:
        fmt.Printf("%d Desember %d\n", tanggal, tahun)
    default:
        fmt.Println("Tidak Valid!!!")
    }

	// Soal 4

	var birthYear = 2002

    var generation string
    if birthYear >= 1944 && birthYear <= 1964 {
        generation = "Baby boomer"
    } else if birthYear >= 1965 && birthYear <= 1979 {
        generation = "Generasi X"
    } else if birthYear >= 1980 && birthYear <= 1994 {
        generation = "Generasi Y (Millennials)"
    } else if birthYear >= 1995 && birthYear <= 2015 {
        generation = "Generasi Z"
    } else {
        generation = "Belum ditentukan nama generasinya"
    }

	fmt.Println("\n====== Jawaban 4 ======")
    fmt.Printf("Waduh!, anda termasuk %s\n", generation)
}