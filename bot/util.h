#pragma once

#include "includes.h"

#define CONVERT_ADDR(x) x & 0xff, (x >> 8) & 0xff, (x >> 16) & 0xff, (x >> 24) & 0xff

void util_strcat(char *, char *);
void util_memcpy(void *, void *, int);
void util_zero(void *, int);

int util_strlen(char *);
int util_strcpy(char *, char *);
int util_atoi(char *, int);
int util_memsearch(char *, int, char *, int);
int util_stristr(char *, int, char *);

BOOL mem_exists(char *, int, char *, int);
BOOL util_strncmp(char *, char *, int);
BOOL util_strcmp(char *, char *);

char *util_itoa(int, int, char *);
char *util_fdgets(char *, int, int);

int util_isupper(char);
int util_isalpha(char);
int util_isspace(char);
int util_isdigit(char);

ipv4_t util_local_addr(void);
