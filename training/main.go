package main

/*
#cgo CFLAGS: -std=gnu99
#include "c/lib.c"
*/
import "C"

import "fmt"

func main() {
	myDC := parse()

	fmt.Print(myDC)
}
