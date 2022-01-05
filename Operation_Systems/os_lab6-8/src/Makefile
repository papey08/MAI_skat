SRC = main.cpp tree.cpp
OBJ = $(SRC:.cpp=.o)

SRC2 = client.cpp timer.cpp
OBJ2 = $(SRC2:.cpp=.o)

all: main client

main: $(OBJ)
	g++ -Wno-unused-variable $(OBJ) -o $@ -lrt -lzmq

client: $(OBJ2)
	g++ -Wno-unused-variable $(OBJ2) -o $@ -lrt -lzmq

.cpp.o:
	g++ -Wno-unused-variable -c $< -o $@

client.o: tree.h timer.h
main.o: tree.h
timer.o: timer.h
tree.o: tree.h

clean:
	rm client.o main.o timer.o tree.o 
