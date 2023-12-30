CC ?= gcc
CFLAGS ?= -Werror -pedantic
main: main.o list.o sort.o
	$(CC) -o main main.o list.o sort.o

main.o: main.c
	$(CC) $(CFLAGS) -c main.c

list.o: list.c list.h
	$(CC) $(CFLAGS) -c list.c

sort.o: sort.c sort.h
	$(CC) $(CFLAGS) -c sort.c

clean:
	rm -rf *.o main
