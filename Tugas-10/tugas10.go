package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

// Soal 1
func showSentence(sentence string, year int) {
	fmt.Printf("sentence: %s\n", sentence)
	fmt.Printf("year: %d\n", year)
}

// Soal 2

func kelilingSegitigaSamaSisi(sisi int, text bool) (string, error) {
	if sisi == 0 {
		err := errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		if text {
			return "", err
		} else {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic recovered:", r)
				}
			}()
			panic(err)
		}
	}

	keliling := sisi * 3
	if text {
		result := fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling)
		return result, nil
	} else {
		return fmt.Sprintf("%d", keliling), nil
	}
}

// Soal 3

func tambahAngka(nilai int, total *int) {
	*total += nilai
}

func cetakAngka(total *int) {
	fmt.Printf("Total angka: %d\n", *total)
}

// Soal 4

var phones = []string{}

func addPhoneData(slice *[]string, data string) {
	*slice = append(*slice, data)
}

// Soal 5

func luasLingkaran(radius float64) float64 {
	luas := math.Pi * math.Pow(radius, 2)
	return math.Round(luas)
}

func kelilingLingkaran(radius float64) float64 {
	keliling := 2 * math.Pi * radius
	return math.Round(keliling)
}

// Soal 6

func luasPersegiPanjang(panjang, lebar float64) float64 {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar float64) float64 {
	return 2 * (panjang + lebar)
}

func main() {
	angka := 1

	// Soal 1

	sentence := "Golang Backend Development"
	year := 2023

	defer showSentence(sentence, year)

	defer fmt.Println("\n====== Soal 1 ======")

	// Soal 2

	fmt.Println("\n====== Soal 2 ======")

	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// Soal 3

	defer cetakAngka(&angka)

	defer fmt.Println("\n====== Soal 3 ======")

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// Soal 4

	fmt.Println("\n====== Soal 4 ======")

	addPhoneData(&phones, "Xiaomi")
	addPhoneData(&phones, "Asus")
	addPhoneData(&phones, "Iphone")
	addPhoneData(&phones, "Samsung")
	addPhoneData(&phones, "Oppo")
	addPhoneData(&phones, "Realme")
	addPhoneData(&phones, "Vivo")

	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second)
	}

	// Soal 5

	fmt.Println("\n====== Soal 5 ======")

	radius := []float64{7, 10, 15}

	for _, r := range radius {
		luas := luasLingkaran(r)
		keliling := kelilingLingkaran(r)

		fmt.Printf("\nJari-jari: %.2f\n", r)
		fmt.Printf("Luas Lingkaran: %.2f\n", luas)
		fmt.Printf("Keliling Lingkaran: %.2f\n", keliling)
	}

	// Soal 6

	fmt.Println("\n====== Soal 6 ======")

	panjang := flag.Float64("panjang", 5, "Panjang persegi panjang")
	lebar := flag.Float64("lebar", 11, "Lebar persegi panjang")
	flag.Parse()

	luas := luasPersegiPanjang(*panjang, *lebar)
	keliling := kelilingPersegiPanjang(*panjang, *lebar)

	fmt.Printf("Panjang: %.2f\n", *panjang)
	fmt.Printf("Lebar: %.2f\n", *lebar)
	fmt.Printf("Luas Persegi Panjang: %.2f\n", luas)
	fmt.Printf("Keliling Persegi Panjang: %.2f\n", keliling)
}

