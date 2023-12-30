#! /bin/bash
echo 'Made by Matvey Popov М80-108Б-20'
rm -rf $2
mkdir $2
cp -r $1 $2
cd $2
cd $1
ls | xargs -ix mv x x_backup
