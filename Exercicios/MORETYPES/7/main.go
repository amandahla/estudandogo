package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			switch {
			case j%15 == 0:
				a[i][j] = 240
			case j%3 == 0:
				a[i][j] = 120
			case j%5 == 0:
				a[i][j] = 150
			default:
				a[i][j] = 100
			}
		}
	}

	return a

}

func main() {
	pic.Show(Pic)
}