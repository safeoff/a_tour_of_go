package main

import "golang.org/x/tour/pic"

func calc(x, y int) uint8 {
	// return uint8((x + y) / 2)
	// return uint8(x * y)
	return uint8(x ^ y)
}

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for y := range p {
		p[y] = make([]uint8, dx)
		for x := range p[y] {
			p[y][x] = calc(x, y)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}

