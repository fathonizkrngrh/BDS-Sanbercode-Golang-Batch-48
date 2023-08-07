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

// Soal 3

func addFruits(fruits *[]string, newFruits ...string) {
	*fruits = append(*fruits, newFruits...)
}

// Soal 4

func addFilms(title, duration, genre, year string, dataFilm *[]map[string]string) {
	newFilm := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}

	*dataFilm = append(*dataFilm, newFilm)
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

	// Soal 3

	fmt.Println("\n====== Soal 3 ======")

	var buah = []string{}

	addFruits(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	for i, fruit := range buah {
		fmt.Printf("%d. %s\n", i+1, fruit)
	}

	// Soal 4

	fmt.Println("\n====== Soal 4 ======")

	var dataFilm = []map[string]string{}

	addFilms("LOTR", "2 jam", "action", "1999", &dataFilm)
	addFilms("avenger", "2 jam", "action", "2019", &dataFilm)
	addFilms("spiderman", "2 jam", "action", "2004", &dataFilm)
	addFilms("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. title : %s\n", i+1, film["title"])
		fmt.Printf("   duration : %s\n", film["duration"])
		fmt.Printf("   genre : %s\n", film["genre"])
		fmt.Printf("   year : %s\n", film["year"])
	}
}