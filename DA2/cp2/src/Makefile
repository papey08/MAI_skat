CXXFLAGS+=-std=c++2a -lstdc++fs -w -O2 -Wextra -Wall -Werror -Wno-sign-compare -Wno-unused-result -pedantic -lm
CXX=g++
FILES=main.cpp LZ77.hpp Huffman.hpp TCompressTreeNode.hpp TCompressTree.hpp TCompressTree.cpp TCompressTree.cpp TCompressTreeNode.cpp TDecompressTree.cpp TDecompressTree.hpp TDecompressTreeNode.cpp TDecompressTreeNode.hpp


run: main.o
	./main.o

main.o: $(FILES)
	$(CXX) $^ $(CXXFLAGS) -o main.o

clean:
	rm -rf main.o

pack:
	zip cp2.zip -r Makefile CMakeLists.txt $(FILES)