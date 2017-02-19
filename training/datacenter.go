package main

type Datacenter struct {
	Rows         int
	Slots        int
	USlotNumber  int
	Pools        int
	ServerNumber int
	Servers      []Server
	USlots       []USlot
	Dcbis        []int
}

type Server struct {
	x, y  int
	Size  int
	Cap   int
	Score int
}

type USlot struct {
	x, y int
}

func NewDatacenter(rows, slots, uslotnumber, pools, servernumber int) *Datacenter {
	return &Datacenter{
		Rows:         rows,
		Slots:        slots,
		USlotNumber:  uslotnumber,
		Pools:        pools,
		ServerNumber: servernumber,
		Servers:      make([]Server, 0, servernumber),
		USlots:       make([]USlot, 0, uslotnumber),
		Dcbis:        make([]int, rows*slots),
	}
}