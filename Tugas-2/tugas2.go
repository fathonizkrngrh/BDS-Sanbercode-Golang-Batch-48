package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Soal 1
	// Tampilkan kalimat "Bootcamp Digital Skill Sanbercode Golang" yang tersusun dari gabungan variabel dalam setiap kata (5 variabel)

	var word1 string = "Bootcamp"
	var word2 string = "Digital"
	var word3 string = "Skill"
	var word4 string = "Sanbercode"
	var word5 string = "Golang"

	sentence := word1 + " " + word2 + " " + word3 + " " + word4 + " " + word5
	fmt.Println("Answer 1 :", sentence)

	// Soal 2
	// Ganti kata "Dunia" menjadi "Golang" menggunakan packages strings

	halo := "Halo Dunia"

	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println("Answer 2 :", halo)

	// Soal 3
	// gabungkan variabel-variabel tersebut agar menghasilkan output
	// saya Senang belajaR GOLANG

	var kataPertama = "saya";
	var kataKedua = "senang";
	var kataKetiga = "belajar";
	var kataKeempat = "golang";

	lastIndexKetiga := len(kataKetiga) - 1

	kataKedua = strings.ToUpper(string(kataKedua[0])) + kataKedua[1:]
	kataKetiga = kataKetiga[:lastIndexKetiga] + strings.ToUpper(string(kataKetiga[lastIndexKetiga]))
	kataKeempat = strings.ToUpper(kataKeempat)

	output := kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat
	fmt.Println("Answer 3 :", output)

	// Soal 4
	// ubahlah variabel diatas menjadi integer dan jumlahkan semuanya

	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	arrayAngka := []string{angkaPertama, angkaKedua, angkaKetiga, angkaKeempat}
	
	total := 0

	for _, angka := range arrayAngka {
		nilai, err := strconv.Atoi(angka)
		if err != nil {
			fmt.Println("Konversi gagal", err)
			return
		}
		total += nilai
	}

	fmt.Println("Answer 4 :", total)

	// Soal 5
	// olah variabel diatas yang hasil output akhrinya adalah "Hi Hi bandung" - 2021

	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = strings.Replace(kalimat, "halo", "Hi", -1)
	angkaStr := strconv.Itoa(angka)

	hasil := `"`+ kalimat + `"` + " - " + angkaStr
	fmt.Println("Answer 5 :",hasil)
}