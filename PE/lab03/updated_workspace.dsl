workspace {
    name "Приложение для хранения файлов"

    model {
        user = person "Пользователь"
        file_storage_service = softwareSystem "file_storage_service" {
            api_service = container "api_service" {
                description "Сервис, который предоставляет API для действий над пользователями, папками и файлами"
                technology "python3/flask"
                tags "service, api"

                api_http_server = component "http_server" {
                    description "Сервер, предоставляющий API для пользователей"
                    technology "flask"
                } 

                api_user_accessor = component "user_accessor" {
                    description "Класс для получения данных пользователя из core_service через nats"
                    technology "nats-py"
                }

                api_core_accessor = component "core_accessor" {
                    description "Класс для получения данных о папках/файлах из core_service через nats"
                    technology "nats-py"
                }

                api_auth_accessor = component "auth_accessor" {
                    description "Класс для обновления/получения токенов, получения id пользователя по access токену"
                    technology "nats-py"
                }

                api_http_server -> api_user_accessor
                api_http_server -> api_core_accessor
                api_http_server -> api_auth_accessor
            }

            core_service = container "core_service" {
                description "Сервис, который управляет файлами/папками (создание, получение, удаление) и правами доступа"
                technology "python3"
                tags "service, core"

                core_nats_server = component "nats_server" {
                    description "Сервер, принимающий сообщения из nats и отвечающий на них"
                    technology "nats-py"
                }

                core_user_accessor = component "user_accessor" {
                    description "Класс для валидации id пользователя через nats"
                    technology "nats-py"
                }

                core_postgres_accessor = component "postgres_accessor" {
                    description "Класс для работы с PostgreSQL"
                    technology "sqlalchemy"
                }

                core_minio_accessor = component "minio_accessor" {
                    description "Класс для работы с minio"
                    technology "minio"
                }

                core_nats_server -> core_user_accessor
                core_nats_server -> core_postgres_accessor
                core_nats_server -> core_minio_accessor
            }

            user_service = container "user_service" {
                description "Сервис, который управляет данными пользователя (фамилия, имя, логин)"
                technology "python3"
                tags "service, user"

                user_nats_server = component "nats_server" {
                    description "Сервер, принимающий сообщения из nats и отвечающий на них"
                    technology "nats-py"
                }

                user_postgres_accessor = component "postgres_accessor" {
                    description "Класс для работы с PostgreSQL"
                    technology "sqlalchemy"
                }

                user_nats_server -> user_postgres_accessor
            }

            auth_service = container "auth_service" {
                description "Сервис, который управляет выдачей токенов доступа (access/refresh)"
                technology "python3, jwt"
                tags "service, auth"

                auth_nats_server = component "nats_server" {
                    description "Сервер, принимающий сообщения из nats и отвечающий на них"
                    technology "nats-py"
                }

                auth_redis_accessor = component "redis_accessor" {
                    description "Класс для работы с Redis"
                    technology "redis"
                }

                auth_nats_server -> auth_redis_accessor
            }

            user_postgres = container "user_postgres" {
                description "База данных PostgreSQL для хранения данных пользователя"
                technology "postgresql"
                tags "db, user"
            }

            core_postgres = container "core_postgres" {
                description "База данных PostgreSQL для хранения файловой структуры пользователя и прав доступа пользователей к файлам и папкам"
                technology "postgresql"
                tags "db, core"
            }

            core_minio = container "core_minio" {
                description "Файловое хранилище minio для хранения содержимого файлов"
                technology "minio"
                tags "storage, core"
            }

            auth_redis = container "auth_redis" {
                description "Хранилище refresh токенов"
                technology "redis"
                tags "db, auth"
            }

            api_user_nats = container "api_user_nats" {
                description "Очередь сообщений для связи api_service и user_service"
                technology "nats"
                tags "nats, api, user"
            }

            api_core_nats = container "api_core_nats" {
                description "Очередь сообщений для связи api_service и core_service"
                technology "nats"
                tags "nats, api, core"
            }

            api_auth_nats = container "api_auth_nats" {
                description "Очередь сообщений для связи api_service и auth_service"
                technology "nats"
                tags "nats, api, auth"
            }

            core_user_nats = container "core_user_nats" {
                description "Очередь сообщений для связи user_service и core_service"
                technology "nats"
                tags "nats, user, core"
            }

            user -> file_storage_service
            user -> api_service
            user -> api_http_server

            api_service -> api_user_nats
            api_user_accessor -> api_user_nats

            api_service -> api_core_nats
            api_core_accessor -> api_core_nats

            api_service -> api_auth_nats
            api_auth_accessor -> api_auth_nats

            core_service -> core_user_nats
            core_user_accessor -> core_user_nats

            api_user_nats -> user_service
            api_user_nats -> user_nats_server

            api_core_nats -> core_service
            api_core_nats -> core_nats_server

            api_auth_nats -> auth_service
            api_auth_nats -> auth_nats_server

            core_user_nats -> user_service
            core_user_nats -> user_nats_server

            user_service -> user_postgres
            user_postgres_accessor -> user_postgres

            auth_service -> auth_redis
            auth_redis_accessor -> auth_redis

            core_service -> core_postgres
            core_postgres_accessor -> core_postgres

            core_service -> core_minio
            core_minio_accessor -> core_minio
        }
    }

    views {
        systemContext file_storage_service "context_view" {
            include *
            autolayout lr
        }

        container file_storage_service "container_view" {
            include *
            autoLayout lr
        }

        component api_service "api_service_view" {
            include *
            autoLayout lr
        }

        component core_service "core_service_view" {
            include *
            autoLayout lr
        }

        component auth_service "auth_service_view" {
            include *
            autoLayout lr
        }

        component user_service "user_service_view" {
            include *
            autoLayout lr
        }

        dynamic file_storage_service "getting_file_usecase" {
            user -> api_service "Пользователь отправляет GET-запрос для поиска файла по названию"
            api_service -> api_core_nats "В очередь сообщений отправляется запрос, который содержит название файла и id пользователя, полученный из jwt токена"
            api_core_nats -> core_service "Брокер сообщений доставляет запрос в core_service"
            core_service -> core_postgres "Выполняется поиск в БД по названию файла и id владельца"
            core_postgres -> core_service "БД возвращает список id из найденных файлов, соответствующих запросу"
            core_service -> api_core_nats "core_service отправляет в очередь сообщений список из ссылок на найденные файлы, составленных по их id"
            api_core_nats -> api_service "Брокер сообщений доставляет ответ от core_service"
            api_service -> user "Сервис возвращает пользователю список ссылок, по которым можно скачать файлы"
        }
    }
}
