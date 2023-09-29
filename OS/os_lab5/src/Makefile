files: main

main: libd1.so libd2.so prog2.cpp
	g++ prog2.cpp -ldl -o main

libd1.so: d1.o 
	g++ -shared d1.o -o libd1.so 

libd2.so: d2.o 
	g++ -shared d2.o -o libd2.so 

d1.o: re1.cpp
	g++ -fPIC -c re1.cpp -o d1.o

d2.o: re2.cpp
	g++ -fPIC -c re2.cpp -o d2.o
