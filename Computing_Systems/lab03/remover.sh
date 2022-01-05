#! /bin/bash

cd $1
for i in $@ 
do
	rm -rf $i
done
