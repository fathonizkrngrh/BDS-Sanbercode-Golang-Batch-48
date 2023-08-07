package main

import (
	"fmt"
	"math"
	"strconv"
)

// Soal 1
func hitungLingkaran(jariJari float64, luasLingkaran *float64, kelilingLingkaran *float64) {
	*luasLingkaran = math.Pi * jariJari * jariJari
	*kelilingLingkaran = 2 * math.Pi * jariJari
}

// Soal 2
func introduce(sentence *string, name, gender, job, age string) {
	var called string

	if gender == "laki-laki" {
		called = "Pak"
	} else {
		called = "Bu"
	}

	ageNum, err := strconv.Atoi(age)
	if err != nil {
		return
	}

	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %d tahun", called, name, job, ageNum)
}


func main() {
	
	// Soal 1

	fmt.Println("====== Soal 1 ======")

	var luasLingkaran float64
	var kelilingLingkaran float64
	
	var jariJariLingkaran float64 = 7
	fmt.Println("Jari-Jari : ", jariJariLingkaran)
	
	hitungLingkaran(jariJariLingkaran, &luasLingkaran, &kelilingLingkaran)

	fmt.Printf("Luas lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling lingkaran: %.2f\n", kelilingLingkaran)
	
	jariJariLingkaran = 12
	fmt.Println("Jari-Jari : ", jariJariLingkaran)
	
	hitungLingkaran(jariJariLingkaran, &luasLingkaran, &kelilingLingkaran)

	fmt.Printf("Luas lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling lingkaran: %.2f\n", kelilingLingkaran)

	// Soal 2

	fmt.Println("\n====== Soal 2 ======")

	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println( sentence) 

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println( sentence)
}