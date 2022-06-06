CXXFLAGS+=-std=c++11 -w -O2 -Wextra -Wall -Werror -Wno-sign-compare -Wno-unused-result -pedantic -lm
CXX=g++



run: main.o
	./main.o

main.o: main.cpp TPair.hpp ISort.hpp
	$(CXX) $^ $(CXXFLAGS) -o main.o

pack:
	zip lab.zip -r Makefile src
