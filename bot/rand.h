#pragma once

#include <stdint.h>

#define PHI 0x9e3779b9

void rand_init(void);
void rand_str(char *, int);

uint32_t rand_next(void);
uint32_t rand_next_range(uint32_t, uint32_t);
