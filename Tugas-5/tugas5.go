package main

import (
	"fmt"
)

// Soal 1

func luasPersegiPanjang(p, l int) int {
    return p * l
}

func kelilingPersegiPanjang(p, l int) int {
    return 2 * (p + l)
}

func volumeBalok(p, l, t int) int {
    return p * l * t
}

// Soal 2

func introduce(name, gender, job, age string) string {
    var called string

    if gender == "laki-laki" {
        called = "Pak "
    } else if gender == "perempuan" {
        called = "Bu "
    } else {
        called = ""
    }

    hasil := fmt.Sprintf("%s%s adalah seorang %s yang berusia %s tahun", called, name, job, age)
    return hasil
}

// Soal 3

func buahFavorit(nama string, buah ...string) string {
    hasil := fmt.Sprintf("Halo nama saya %s dan buah favorit saya adalah ", nama)

    if len(buah) > 0 {
        hasil += fmt.Sprintf(`"%s"`, buah[0])
        for i := 1; i < len(buah); i++ {
            hasil += fmt.Sprintf(`, "%s"`, buah[i])
        }
    } else {
        hasil += "tidak punya buah favorit"
    }

    return hasil
}

// Soal 4

var dataFilm = []map[string]string{}

func tambahDataFilm() func(string, string, string, string) {
	return func(title, jam, genre, tahun string) {
		dataFilm = append(dataFilm, map[string]string{
			"title": title,
			"jam":   jam,
			"genre": genre,
			"tahun": tahun,
		})
	}
}


func main() {
	// Soal 1
	fmt.Println("====== Soal 1 ======")

	panjang := 12
    lebar := 4
    tinggi := 8

    luas := luasPersegiPanjang(panjang, lebar)
    keliling := kelilingPersegiPanjang(panjang, lebar)
    volume := volumeBalok(panjang, lebar, tinggi)

    fmt.Println("Luas Persegi Panjang:", luas)
    fmt.Println("Keliling Persegi Panjang:", keliling)
    fmt.Println("Volume Balok:", volume)

	// Soal 2

	fmt.Println("\n====== Soal 2 ======")

	john := introduce("John", "laki-laki", "penulis", "30")
    fmt.Println(john) 

    sarah := introduce("Sarah", "perempuan", "model", "28")
    fmt.Println(sarah) 

	// Soal 3

	fmt.Println("\n====== Soal 3 ======")

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
    buahFavoritJohn := buahFavorit("John", buah...)

    fmt.Println(buahFavoritJohn)

	// Soal 4

	fmt.Println("\n====== Soal 4 ======")

	tambahDataFilm := tambahDataFilm()

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}