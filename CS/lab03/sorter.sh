#! /bin/bash
cd $1
rm -rf output.txt
sort -t : -k2 $2 > output.txt
