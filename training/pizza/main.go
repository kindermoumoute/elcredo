package main

import (
	"fmt"
	"math/rand"
)

func main() {
	myPizza := parse("input/small.in")

	myPizza.cutPizza()

	write(myPizza.Slice)
}

func (p *Pizza) cutPizza() {
	sizes := getSizes(p.MaxCells, p.Rows)
	tempCells := make(map[int]*Cell)
	for i, cell := range p.Cell {
		tempCells[i] = &cell
	}
	for len(tempCells) > 0 {
		p.print(tempCells)
		cellIndex := rand.Int() % len(tempCells)
		fmt.Println(cellIndex, len(tempCells))
		bestSlice := p.findBestSlice(cellIndex, sizes, tempCells)
		if bestSlice == nil {
			fmt.Println("deleted : ", cellIndex)
			delete(tempCells, cellIndex)
			continue
		}
		p.Slice = append(p.Slice, bestSlice)
		p.removeCells(tempCells, bestSlice)
	}
}

func (p *Pizza) removeCells(m map[int]*Cell, s *Slice) {
	for x := s.X1; x < s.X2; x++ {
		for y := s.Y1; y < s.Y2; y++ {
			fmt.Println("deleted from slice : ", y*p.Rows+x)
			delete(m, y*p.Rows+x)
		}
	}
}

func (p *Pizza) print(m map[int]*Cell) {
	fmt.Printf("\n\nPIZZA INITIALE")
	letter := []string{"T", "M"}
	for i, c := range p.Cell {
		if i%p.Rows == 0 {
			fmt.Printf("\n")
		}
		fmt.Print(letter[c.Ingredient])
	}
	fmt.Printf("\nPIZZA DECOUPEE")
	for i, c := range p.Cell {
		if i%p.Rows == 0 {
			fmt.Printf("\n")
		}
		_, unused := m[i]
		if !unused {
			fmt.Print(letter[c.Ingredient])

		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("\n")
}
func (p *Pizza) findBestSlice(cellIndex int, sizes []size, m map[int]*Cell) *Slice {
	var i int
	for _, s := range sizes {
		var LTomato, LMushroom int
		brk := false
		inc := []int{1, 1, 1, -1, -1, -1, -1, 1}
		for incIndex := 0; incIndex < 4; incIndex++ {
			incx := inc[incIndex*2]
			for x := 0; x < s.row*incx; x += incx {
				incy := inc[incIndex*2+1]
				for y := 0; y < s.column*incy; y += incy {
					i = cellIndex + (y * p.Rows) + x
					// break if the cell is already used
					if i < 0 || i >= len(p.Cell) || p.Cell[i].Used != nil {
						brk = true
						break
					}
					if p.Cell[i].Ingredient == TOMATO {
						LTomato++
					} else {
						LMushroom++
					}
				}
				if brk {
					break
				}
			}
			if brk {
				break
			}
			if LMushroom >= p.MinIngredient && LTomato >= p.MinIngredient {
				X1 := cellIndex % p.Rows
				X2 := cellIndex%p.Rows + s.row
				Y1 := cellIndex / p.Rows
				Y2 := cellIndex/p.Rows + s.column
				if X1 > X2 {
					X1, X2 = X2, X1
				}
				if Y1 > Y2 {
					Y1, Y2 = Y2, Y1
				}
				return &Slice{X1: X1, X2: X2, Y1: Y1, Y2: Y2}
			}
		}
	}
	return nil
}
