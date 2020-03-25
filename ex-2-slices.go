package main

import "image"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		res[i] = make([]uint8, dx)
	}
	return res
}

func main() {
	pic.Show(Pic)
}
