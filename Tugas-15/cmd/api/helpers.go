package main

import (
	"strconv"
)

func GetIndeksNilai(nilai int) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 {
		return "B"
	} else if nilai >= 60 {
		return "C"
	} else if nilai >= 50 {
		return "D"
	}
	return "E"
}

func ParseInt(s string) int {
	val, _ := strconv.ParseInt(s, 10, 32)
	return int(val)
}

