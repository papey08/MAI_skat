#! ?bin?bash
echo 'Made by Matvey Popov М80-108Б-20'
read path
cd $path
ls -p | grep -v "/"| sort -r| xargs wc -l
