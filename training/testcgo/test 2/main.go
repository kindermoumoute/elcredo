package main

/*
#cgo CFLAGS: -std=gnu99
#include "sources/lib.c"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type slot struct {
	x, y int
	tab []int
}

func main() {
	a := slot{x:2, y:3, tab: make([]int, 12)}
	for i := range a.tab {
		a.tab[i] = i
	}

	C.add(unsafe.Pointer(&a))
	fmt.Println(a)
}

// GODEBUG=cgocheck=0 go run main.go
