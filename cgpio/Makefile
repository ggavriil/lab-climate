IDIR=.
CC=gcc
CFLAGS=-I$(IDIR) -Wall -Wextra -std=gnu99 

ODIR=obj

LDIR=.

LIBS=

_DEPS=gpio.h
DEPS=$(patsubst %,$(IDIR)/%,$(_DEPS))

_OBJ=gpio.o
OBJ=$(patsubst %,$(ODIR)/%,$(_OBJ))

$(ODIR)/%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

libgpio: $(OBJ)
	ar rcs ../lib/libgpio.a $a $^

.PHONY: clean

clean: 
	rm -f $(ODIR)/*.o *~ core $(INCDIR)/*~
