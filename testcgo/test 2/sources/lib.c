#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <inttypes.h>

typedef struct{
	int64_t *tab;
	int64_t len;
	int64_t cap;
}SLICE;

typedef struct{
	int64_t x, y;
	SLICE slice;
}SLOT;

void add(void *a) {
	SLOT *pSlot = (SLOT*)a;
	SLOT slot = *(SLOT*)a;

	for(int64_t i = 0; i < slot.slice.len; i++)
		slot.slice.tab[i] = 2*i;
}

