package main

import "fmt"

// Soal 1

type Buah struct {
    Nama      string
    Warna     string
    AdaBijinya bool
    Harga     int
}

// Soal 2

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

// Soal 3

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColor(newColor string) {
	p.colors = append(p.colors, newColor)
}

// Soal 4

type movie struct {
	title   string
	genre   string
	duration int
	year    int
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	newMovie := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}
	*dataFilm = append(*dataFilm, newMovie)
}

func main() {

	// Soal 1

	fmt.Println("====== Soal 1 ======")

    buahData := []Buah{
        {Nama: "Nanas", Warna: "Kuning", AdaBijinya: false, Harga: 9000},
        {Nama: "Jeruk", Warna: "Oranye", AdaBijinya: true, Harga: 8000},
        {Nama: "Semangka", Warna: "Hijau & Merah", AdaBijinya: true, Harga: 10000},
        {Nama: "Pisang", Warna: "Kuning", AdaBijinya: false, Harga: 5000},
    }

    for _, buah := range buahData {
        fmt.Printf("Nama: %s\nWarna: %s\nAda Bijinya: %v\nHarga: %d\n\n",
            buah.Nama, buah.Warna, buah.AdaBijinya, buah.Harga)
    }

	// Soal 2

	fmt.Println("\n====== Soal 2 ======")

	segitiga1 := segitiga{alas: 6, tinggi: 8}
	persegi1 := persegi{sisi: 5}
	persegiPanjang1 := persegiPanjang{panjang: 7, lebar: 4}

	fmt.Println("Luas Segitiga:", segitiga1.luas())
	fmt.Println("Luas Persegi:", persegi1.luas())
	fmt.Println("Luas Persegi Panjang:", persegiPanjang1.luas())

	// Soal 3

	fmt.Println("\n====== Soal 3 ======")

	iphone := phone{
		name:   "iPhone 13",
		brand:  "Apple",
		year:   2022,
		colors: []string{"Black", "White"},
	}

	fmt.Println("Sebelum :", iphone)

	iphone.addColor("Blue")
	iphone.addColor("Red")

	fmt.Println("Sesudah :", iphone)

	// Soal 3

	fmt.Println("\n====== Soal 4 ======")

	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. title : %s\n", i+1, film.title)
		fmt.Printf("   duration : %d jam\n", film.duration)
		fmt.Printf("   genre : %s\n", film.genre)
		fmt.Printf("   year : %d\n\n", film.year)
	}
}