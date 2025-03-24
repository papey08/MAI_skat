# lab03

## line_counter

Входные данные: путь к директории

Выходные данные: утилита должна печатать для каждого файла в данном каталогое: 
1) его имя
2) количество строк в данном файле

Вывод должен быть отсортирован по лексиграфическому убывание

Пример:
```
nikitatrubcenko@MacBook-Pro-Nikita var1 % ./lines_counter.sh .
      11 total
       7 task1
       4 lines_counter.sh
```

## remover

Входные данные: путь к директории и список файлов, которые в данной директории нужно удалить

Необходимо удалить заданные файлы в заданной директории

## sorter

Входные данные: путь к файлу, пары число-строка, разделенные знаком :
Выходные данные: файл output.txt, который содержит файлы из input.txt, но в 
отсортированном порядке по убыванию строк. 

Пример:
```
nikitatrubcenko@MacBook-Pro-Nikita var3 % cat input.txt 
1:fdkgdfkj
2:krdgkrdlgk
3:rkdgrdg
4:krgkdrgmr
5:gdhdhgdf
6:gfkdgmkdmg
7:tkfhkfmhkftmhk
8:aaaawdwa
nikitatrubcenko@MacBook-Pro-Nikita var3 % ./sorter.sh output.txt 
nikitatrubcenko@MacBook-Pro-Nikita var3 % cat output.txt 
8:aaaawdwa
1:fdkgdfkj
5:gdhdhgdf
6:gfkdgmkdmg
2:krdgkrdlgk
4:krgkdrgmr
3:rkdgrdg
7:tkfhkfmhkftmhk
```

## backuper

Входные данные: ./backuper.sh [path_to_dir] [path_to_reserve_copy]

Выходные данные: резервная копия директории, находящейся по пути [path_to_dir]. 
Резервная копия должна находится в [path_to_reserve_copy]

К имени каждого файла в резервной копии должен быть добавлен суффикс _backup

## remote_copier

Входные данные: ./remote_copier [username] [remote_server_ip] [local_path]

Выходные данные: утилита должна создать архив из директории, находящейся по 
пути [local_path], сжать данный архив и скопировать его на удаленный сервер по 
адресу [remote_server_ip] в папку ~. [username] -- имя пользователя на 
удаленном сервере. 
