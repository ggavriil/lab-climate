#ifndef GPIO_INCLUDED
#define GPIO_INCLUDED

#include <stdbool.h>

#define GPIO_BASE 0x20200000
#define BLOCK_SIZE (4 * 1024)
#define PAGE_SIZE (4 * 1024)

void input_pin(int p);

void output_pin(int p);

bool is_set(int p);

void set_pin(int p);

void clear_pin(int p);

void write_pin(int p, bool v);

void init_gpio(void);




#endif
