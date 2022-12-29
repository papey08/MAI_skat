CXXFLAGS+=-std=c++14 -w -O2 -Wextra -Wall -Werror -Wno-sign-compare -Wno-unused-result -pedantic -lm
CXX=g++
FILES=main.cpp TSuffixTree.hpp


run: main.o
	./main.o

main.o: $(FILES)
	$(CXX) $^ $(CXXFLAGS) -o main.o

clean:
	rm -rf main.o

pack:
	zip solve.zip -r Makefile $(FILES)
