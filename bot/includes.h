#pragma once

#include <unistd.h>
#include <stdint.h>
#include <stdarg.h>

#define STDIN   0
#define STDOUT  1
#define STDERR  2

#define FALSE   0
#define TRUE    1
typedef char BOOL;

typedef uint32_t ipv4_t;
typedef uint16_t port_t;

#define INET_ADDR(o1,o2,o3,o4) (htonl((o1 << 24) | (o2 << 16) | (o3 << 8) | (o4 << 0)))

#define CNC_PORT 38241
#define SINGLE_INSTANCE_PORT 8345

#define CNC_OP_PING         0x00
#define CNC_OP_KILLSELF     0x10
#define CNC_OP_KILLATTKS    0x20
#define CNC_OP_PROXY        0x30
#define CNC_OP_ATTACK       0x40

extern ipv4_t LOCAL_ADDR;
