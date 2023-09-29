#!/bin/bash
pushd .
path=`pwd`
echo 'Current path is '$path
echo 'Files and directories in '$path
ls
rm -rf dir1
mkdir dir1
cd dir1
> txt1.txt
> txt2.txt
> txt3.txt
echo 'Message from txt1.txt' > txt1.txt
echo 'Message from txt2.txt' > txt2.txt
cat txt2.txt > txt3.txt
rm -rf dir2
mkdir dir2
mv txt3.txt dir2
cp txt1.txt dir2
cp txt2.txt dir2
cd dir2
find txt2.txt | xargs rm
mv txt3.txt new_txt2.txt
popd
echo 'New files and directories in '$path
ls
