package main

import "math"

type size struct {
	row, column int
}

func getSizes(h, max int) []size {
	sqrt := math.Sqrt(float64(h))
	sizes := make([]size, 0, int(sqrt))
	for i := 1; float64(i) <= sqrt; i++ {
		if h%i == 0 && i <= max {
			sizes = append(sizes, size{row: i, column: h / i})
		}
	}
	return sizes
}

//func (p *Pizza) splitPizza() {
//	sizes := getSizes(p.MaxCells, p.Rows)
//	for _, s := range sizes {
//		for x := 0; x < p.Rows; x++ {
//			for y := 0; y < p.Columns; y++ {
//				if placerSlide(p, &Slice{X1: x, Y1: y, X2: x + s.row, Y2: y + s.column}) {
//					y = y + s.column
//				}
//			}
//		}
//	}
//}
