package main

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"tugas11/balok"
	"tugas11/tabung"
)

// Soal 2

func getMovies(ch chan<- string, movies ...string) {
	for i, movie := range movies {
		ch <- fmt.Sprintf("%d. %s\n", i+1, movie)
	}
	close(ch)
}


func main() {
  // Soal 1

  fmt.Println("\n====== Soal 1 ======")

	phones := []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	var wg sync.WaitGroup
	wg.Add(len(phones))

  sort.Strings(phones)

	for idx, phone := range phones {
		go func(index int, p string) {
			defer wg.Done()
			time.Sleep(time.Duration(index) * time.Second) // Tunggu sesuai indeks
			fmt.Printf("%d. %s\n", index+1, p)
		}(idx, phone)
	}

	wg.Wait()

  // Soal 2

  fmt.Println("\n====== Soal 2 ======")

  movies := []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for title := range moviesChannel {
    fmt.Printf(title)
	}

  // Soal 3

  fmt.Println("\n====== Soal 3 ======")

  radii := []float64{8, 14, 20}
	tinggi := 10

	luasCh := make(chan float64)
	kelilingCh := make(chan float64)
	volumeCh := make(chan float64)

	for _, radius := range radii {
		go tabung.LuasLingkaran(radius, luasCh)
		go tabung.KelilingLingkaran(radius, kelilingCh)
		go tabung.VolumeTabung(radius, float64(tinggi), volumeCh)
	}

	fmt.Println("Hasil")
	for _, radius := range radii {
		luas := <-luasCh
		keliling := <-kelilingCh
		volume := <-volumeCh

		fmt.Printf("Jari-jari %.0f:\n", radius)
		fmt.Printf("Luas Lingkaran: %.2f\n", luas)
		fmt.Printf("Keliling Lingkaran: %.2f\n", keliling)
		fmt.Printf("Volume Tabung: %.2f\n\n", volume)
	}

	// Soal 4

	fmt.Println("====== Soal 4 ======")

	length := 5.0
	width := 3.0
	height := 4.0

	luasPersegiPanjangCh := make(chan float64)
	kelilingPersegiPanjangCh := make(chan float64)
	volumeBalokCh := make(chan float64)

	go balok.LuasPersegiPanjang(length, width, luasPersegiPanjangCh)
	go balok.KelilingPersegiPanjang(length, width, kelilingPersegiPanjangCh)
	go balok.VolumeBalok(length, width, height, volumeBalokCh)

	for i := 0; i < 3; i++ {
		select {
		case area := <-luasPersegiPanjangCh:
			fmt.Printf("Luas Persegi Panjang: %.2f\n", area)
		case perimeter := <-kelilingPersegiPanjangCh:
			fmt.Printf("Keliling Persegi Panjang: %.2f\n", perimeter)
		case volume := <-volumeBalokCh:
			fmt.Printf("Volume Persegi Panjang: %.2f\n", volume)
		}
	}
}