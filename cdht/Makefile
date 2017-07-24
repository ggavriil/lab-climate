IDIR=../cgpio
CC=gcc
CFLAGS=-I$(IDIR) -Wall -Wextra -std=gnu99

ODIR=obj

LDIR=.

LIBS=

_DEPS=gpio.h
DEPS=$(patsubst %,$(IDIR)/%,$(_DEPS))

_OBJ=common_dht_read.o pi_dht_read.o
OBJ=$(patsubst %,$(ODIR)/%,$(_OBJ))

$(ODIR)/%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)



libdht: $(OBJ)
	ar rcs ../lib/libdht.a $a $^

#dht: $(OBJ)
#	gcc -o $@ $^ $(CFLAGS) $(LIBS)

.PHONY: clean

clean: 
	rm -f $(ODIR)/*.o *~ core $(INCDIR)/*~
