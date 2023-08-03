package main

import "fmt"

func main() {
	// Soal 1

	fmt.Println("====== Soal 1 ======")

	for i := 1; i <= 30; i++ {
		if i%2 == 0 {
			fmt.Println(i, "- Berkualitas")
		} else if i%3 == 0 {
			fmt.Println(i, "- I Love Coding")
		} else {
			fmt.Println(i, "- Santai")
		}
	}

	// Soal 2

	fmt.Println("\n====== Soal 2 ======")

	tinggi := 7

	for i := 1; i <= tinggi; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("#")
        }
        fmt.Println()
    }

	// Soal 3

	fmt.Println("\n====== Soal 3 ======")

	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

    output := kalimat[2:7]

    fmt.Println(output)

	// Soal 4

	fmt.Println("\n====== Soal 4 ======")

	var sayuran = []string{}

    sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

    for i, sayur := range sayuran {
        fmt.Printf("%x. %s\n", i+1, sayur)
    }

	// Soal 5

	fmt.Println("\n====== Soal 5 ======")

	var satuan = map[string]int{
        "panjang": 7,
        "lebar":   4,
        "tinggi":  6,
    }

	var volume = 1

	for s, v := range satuan {
        fmt.Printf("%s = %2d\n", s, v)
		volume *= v
    }

    fmt.Printf("volume balok = %d\n", volume)
}