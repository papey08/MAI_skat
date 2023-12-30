#! /bin/bash
echo 'Made by Matvey Popov М80-108Б-20'
tar -cf $3.gz $3
scp -p 6789 $3.gz $1@$2:/stud/$1
