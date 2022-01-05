#!/bin/bash
rm -rf testdir
mkdir testdir
cd testdir
rm -rf text1.txt
rm -rf text2.txt
> text1.txt
> text2.txt
echo 'aaa aab aba abb baa bab bba bbb' > text1.txt
echo 'aaa aac aca acc caa cac cca ccc' > text2.txt
echo '*diff* Files text1.txt & text2.txt:'
diff -y text1.txt text2.txt
echo '*cmp* Differences between text1.txt & text2.txt:'
cmp -l text1.txt text2.txt
echo '*wc* Number of words in text1.txt:'
wc -w text1.txt
echo '*tail* Last 4 bytes of text2.txt:'
tail -c 4 text2.txt
echo '*head* First 8 bytes of text1.txt:'
head -c 8 text1.txt
echo '*du* Size of text1.txt:'
du -h text1.txt
echo '*cut* Every second byte of text1.txt'
cut -b 2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32 text1.txt
rm -rf reserve.txt
> reserve.txt
dd if=text2.txt of=reserve.txt
echo '*dd* reserve.txt (copy of text2.txt):'
cat reserve.txt
echo '*gzip* Making reserve.txt.gz from reserve.txt'
gzip reserve.txt
echo '*xargs* Removing text1.txt & text2.txt'
ls *text* | xargs rm
echo '*gzip* Making reserve.txt from reserve.txt.gz'
gunzip reserve.txt.gz
echo 'reserve.txt:'
cat reserve.txt
cd ..
rm -r testdir