package tabung

import "math"

func LuasLingkaran(r float64, ch chan<- float64) {
	area := math.Pi * r * r
	ch <- area
}

func KelilingLingkaran(r float64, ch chan<- float64) {
	circumference := 2 * math.Pi * r
	ch <- circumference
}

func VolumeTabung(r, t float64, ch chan<- float64) {
	volume := math.Pi * r * r * t
	ch <- volume
}