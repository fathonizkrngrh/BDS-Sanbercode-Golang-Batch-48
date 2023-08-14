package balok

func LuasPersegiPanjang(p, l float64, ch chan<- float64) {
	area := p * l
	ch <- area
}

func KelilingPersegiPanjang(p, l float64, ch chan<- float64) {
	perimeter := 2 * (p + l)
	ch <- perimeter
}

func VolumeBalok(p, l, t float64, ch chan<- float64) {
	volume := p * l * t
	ch <- volume
}
