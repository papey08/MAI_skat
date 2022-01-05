#!/bin/bash

case $1 in
   	"--help")
		echo "Скрипт способен заменять разделители пути у файлов с Makefile в названии
Флаги:
--a заменяем / на \ 
--b заменяем \ на /
--ac dir переходим по директории dir, заменяем / на \ 
--bc dir переходим по директории dir, заменяем \ на / 
--test1 создаёт первый тест
--test2 создаёт второй тест
--test3 создаёт третий тест
--rmtest удаляет тест

Script is able to replace separator of path in files with Makefile in the name
Flags:
--a replace / on \ 
--b replace \ on / 
--ac dir go to directory dir, replace / on \ 
--bc dir go to directory dir, replace \ on / 
--test1 creates the first test
--test2 creates the second test
--test3 creates the third test
--rmtest removes test"
	;;
	"--a")
		dir=`find -name "Makefile*"`
		echo ${dir////\\}
	;;
	"--b")
		dir=`find -name "Makefile*"`
		echo ${dir//\///}
	;;
	"--ac")
		cd $2
		dir=`find -name "Makefile*"`
		echo ${dir////\\}
	;;
	"--bc")
		cd $2
		dir=`find -name "Makefile*"`
		echo ${dir//\///}
	;;
		"--test1")
		rm -rf test
		rm -f file_Makefile.txt
		rm -f Makefile_file.txt
		rm -f makefile.txt
		mkdir test
		cd test
		>file_Makefile.txt
		>Makefile_file.txt
		>makefile.txt
	;;
	"--test2")
		rm -Rf test
		rm -f file_Makefile.txt
		rm -f Makefile_file.txt
		rm -f makefile.txt
		mkdir test
		cd test
		>Makefile1.txt
		mkdir dir1
		mkdir dir2
		cd dir1
		>Makefile2.txt
		cd ..
		cd dir2
		>file_Makefile.txt
	;;
	"--test3")
		rm -Rf test
		rm -f file_Makefile.txt
		rm -f Makefile_file.txt
		rm -f makefile.txt
		>file_Makefile.txt
		>Makefile_file.txt
		>makefile.txt
	;;
	"--rmtest")
		rm -Rf test
		rm -f file_Makefile.txt
		rm -f Makefile_file.txt
		rm -f makefile.txt
	;;
esac