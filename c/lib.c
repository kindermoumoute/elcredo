#include <stdio.h>
#include <stdlib.h>
#include <math.h>

typedef struct {
	int x;
	int y;
	int size;
	int cap; // CAPACITY OF THE SERVER
	int score;
} SERVER;

typedef struct {
	int x;
	int y;
} USLOT;

typedef struct{
	SERVER *tab;
	int64_t len;
	int64_t cap;
}SLICESERVER;

typedef struct{
	USLOT *tab;
	int64_t len;
	int64_t cap;
}SLICEUSLOT;

typedef struct {
	int Rows;
	int Slots;
	int USlotNumber;
	int Pools;
	int ServerNumber;
	SLICESERVER Servers;
	SLICEUSLOT USlots;
} DATACENTER;

DATACENTER *dc;

int calculateCoord(int x, int y, int xMax) {
	return xMax * y + x;
}

void fillDcBisUSlot(int *dcBis) {
	int nbrUSlots;
	int coord;
	USLOT *pUSlot;

	for(nbrUSlots = 0; nbrUSlots < *(dc->USlotNumber); nbrUSlots++) {
		pUSlot = dc->USlots[nbrUSlots];
		coord = calculateCoord(*pUSlot->x, *pUSlot->y, *dc->Slots);
		*dcBis[coord] = 1;
	}
}

void sortDatacenter(void *pDc, int *dcBis) {
	dc = (DATACENTER*)pDc;

	fillDcBisUSlot(dcBis);
}




