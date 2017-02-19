package main

type Pizza struct {
	Rows, Columns int
	MinIngredient int
	MaxCells      int
	Cell          []Cell
}

type Cell struct {
	Ingredient int
	Used       *Slice
}

type Slice struct {
	X1, X2 int
	Y1, Y2 int
}

const (
	TOMATO = iota
	MUSHROOM
)

func NewPizza(rows, columns, miningredient, maxcells int) Pizza {
	return Pizza{Rows: rows, Columns: columns, MinIngredient: miningredient, MaxCells: maxcells, Cell: make([]Cell, rows*columns)}
}
