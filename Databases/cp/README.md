# Курсовой проект

## Вариант 29. Создание и заполнение отношений БД фитнес-клуба

Вместе со мной над проектом работали [artemmoroz0v](https://github.com/artemmoroz0v) и [silverfatt](https://github.com/silverfatt)

## Подробное описание проекта есть в [отчёте](https://github.com/papey08/MAI_skat/blob/main/Databases/cp/docs/report.pdf)

## Запуск с докером

* В этой директории выполнить команду 

```bash
docker-compose up
```

* В своём любимом браузере переходим по адресу localhost:8080
* Готово!

## Запуск без докера

* Установить Go 1.19
* Установить pgadmin
* В pgadmin создать пустую базу данных, в консоли выполнить файл *DB_init.sql*
* В файле *config.yml* вставить свои настройки БД
* В терминале в директории проекта выполнить команды

```(Go)
go env -w GO111MODULE=on
go mod download
go run main.go
```

* В своём любимом браузере переходим по адресу localhost:8080
* Готово!
