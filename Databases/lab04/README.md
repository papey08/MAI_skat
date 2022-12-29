# lab04

## Для запуска

* Установить Go 1.19
* Установить pgadmin
* В pgadmin создать пустую базу данных, в консоли выполнить файл *DB_init.sql*
* В файле *main.go* вставить настройки БД в строку:

```(Go)
err := server.OpenDB("user=user password=password dbname=dbname sslmode=disable")
```

* В терминале в директории проекта выполнить команды

```(Go)
go env -w GO111MODULE=on
go mod download
go run main.go
```

* В своём любимом браузере переходим по адресу localhost:8080
* Готово!

## Докера нет и не будет!!
