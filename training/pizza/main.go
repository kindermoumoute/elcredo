package main

import "fmt"

func main() {
	myPizza := parse("input/small.in")
    
    fmt.Println(myPizza)
    
    mySlice := &Slice{X1:0, X2:4, Y1:0, Y2:0}
    
    fmt.Println(sliceIsValid(myPizza, mySlice))
}
