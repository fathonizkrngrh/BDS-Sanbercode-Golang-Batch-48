package phoneinfo

import (
	"fmt"
)

type Phone struct {
	Name   string
	Brand  string
	Year   int
	Colors []string
}

type DisplayInfo interface {
	GetInfo() string
}

func (p Phone) GetInfo() string {
	info := fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\nColors: %v", p.Name, p.Brand, p.Year, p.Colors)
	return info
}
