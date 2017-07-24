#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <sys/mman.h>
#include <stdbool.h>
#include <unistd.h>
#include "gpio.h"

#define STOP while(1){}

volatile unsigned *gpio;

void input_pin(int p) {
    *(gpio + p / 10) &= ~(7 << ((p % 10) * 3));
}

void output_pin(int p) {
    *(gpio + p / 10) |= (1 << ((p % 10) * 3));
}

bool is_set(int p) {
    return *(gpio + 13) & (1 << p);
}

void set_pin(int p) {
    *(gpio + 7) = 1 << p;
}

void clear_pin(int p) {
    *(gpio + 10) = 1 << p;
}

void write_pin(int p, bool v) {
    if(v) {
        set_pin(p);
    } else {
        clear_pin(p);
    }
}

void init_gpio(void) {
    int mem_fd;
    void* gpio_map;
    if((mem_fd = open("/dev/gpiomem", O_RDWR|O_SYNC)) < 0) {
        printf("can't open /dev/mem \n");
        exit(EXIT_FAILURE);
    }
    gpio_map = mmap(NULL, BLOCK_SIZE, PROT_READ|PROT_WRITE, MAP_SHARED, mem_fd, GPIO_BASE);
    close(mem_fd);
    if(gpio_map == MAP_FAILED) {
        printf("mmap error %d\n", (int) gpio_map);
        exit(EXIT_FAILURE);
    }
    gpio = (volatile unsigned*) gpio_map;
}

/*
int main(void) {
    init_gpio();
    input_pin(17);
    output_pin(17);
    while(1) {
        set_gpio(17);
        sleep(1);
        clear_gpio(17);
        sleep(1);
    }
    return 0;
}
*/
