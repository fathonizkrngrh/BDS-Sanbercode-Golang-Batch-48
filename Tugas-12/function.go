package main

import "math"

func LuasLingkaran(jariJari float64) float64 {
	return math.Pi * math.Pow(jariJari, 2)
}

func KelilingLingkaran(jariJari float64) float64 {
	return 2 * math.Pi * jariJari
}

func VolumeTabung(jariJari, tinggi float64) float64 {
	return math.Pi * math.Pow(jariJari, 2) * tinggi
}