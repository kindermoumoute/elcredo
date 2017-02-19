package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parse(path string) *Pizza {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	tmp := bytes.NewReader(dat)
	s := bufio.NewScanner(tmp)
	s.Scan()
	args := toSliceOfInt(s.Text())
	myPizza := NewPizza(args[0], args[1], args[2], args[3])
	i := 0
	for s.Scan() {
		for _, c := range s.Bytes() {
			switch c {
			case byte('T'):
				myPizza.Cell[i].Ingredient = TOMATO
			case byte('M'):
				myPizza.Cell[i].Ingredient = MUSHROOM
			}
			i++
		}
	}

	return &myPizza
}

func encode(slices []*Slice) string {
	output := fmt.Sprintf("%d", len(slices))
	for _, slice := range slices {
		output += fmt.Sprintf("\n%d %d %d %d", slice.X1, slice.Y1, slice.X2, slice.Y2)
	}
	return output
}
func write(slices []*Slice) {
	output := encode(slices)
	err := ioutil.WriteFile("output/out.out", []byte(output), 0644)
	if err != nil {
		return
	}
}
func toSliceOfInt(line string) []int {
	args := strings.Split(line, " ")
	rep := make([]int, len(args))
	var err error
	for i, v := range args {
		rep[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}
	return rep
}
