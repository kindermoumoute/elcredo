package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
)

func parse() *Datacenter {
	dat, err := ioutil.ReadFile("entry/dc.in")
	if err != nil {
		panic(err)
	}
	tmp := bytes.NewReader(dat)
	s := bufio.NewScanner(tmp)
	s.Scan()
	args := toSliceOfInt(s.Text())
	myDC := NewDatacenter(args[0], args[1], args[2], args[3], args[4])

	uslotNb := args[2]
	for s.Scan() {
		if uslotNb > 0 {
			args := toSliceOfInt(s.Text())
			myDC.USlots = append(myDC.USlots, USlot{x: args[0], y: args[1]})
			uslotNb--
		} else {
			args := toSliceOfInt(s.Text())
			myDC.Servers = append(myDC.Servers, Server{Size: args[0], Cap: args[1]})
		}
	}
	return myDC
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
