workspace {
    name "Приложение для хранения файлов"

    model {
        user = person "Пользователь"
        file_storage_service = softwareSystem "file_storage_service" {
            api_service = container "api_service" {
                description "Сервис, который предоставляет API для действий над пользователями, папками и файлами"
                technology "python3/fastapi"
                tags "service, api"

                api_http_server = component "http_server" {
                    description "Сервер, предоставляющий API для пользователей"
                    technology "fastapi"
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
                description "Сервис, который управляет файлами/папками (создание, получение, удаление)"
                technology "python3"
                tags "service, core"

                core_nats_server = component "nats_server" {
                    description "Сервер, принимающий сообщения из nats и отвечающий на них"
                    technology "nats-py"
                }
                
                core_memory_storage = component "memory_storage" {
                    description "Класс для хранения папок и файлов пользователей в памяти"
                }

                core_nats_server -> core_memory_storage
            }

            user_service = container "user_service" {
                description "Сервис, который управляет данными пользователя (фамилия, имя, логин)"
                technology "python3"
                tags "service, user"

                user_nats_server = component "nats_server" {
                    description "Сервер, принимающий сообщения из nats и отвечающий на них"
                    technology "nats-py"
                }

                user_memory_accessor = component "memory_accessor" {
                    description "Класс для хранения пользователей в памяти"
                }

                user_nats_server -> user_memory_accessor
            }

            auth_service = container "auth_service" {
                description "Сервис, который управляет выдачей токенов доступа (access/refresh)"
                technology "python3, jwt"
                tags "service, auth"

                auth_nats_server = component "nats_server" {
                    description "Сервер, принимающий сообщения из nats и отвечающий на них"
                    technology "nats-py"
                }

                auth_memory_accessor = component "memory_accessor" {
                    description "Класс для хранения refresh-токенов в памяти"
                }

                auth_nats_server -> auth_memory_accessor
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

            user -> file_storage_service
            user -> api_service
            user -> api_http_server

            api_service -> api_user_nats
            api_user_accessor -> api_user_nats

            api_service -> api_core_nats
            api_core_accessor -> api_core_nats

            api_service -> api_auth_nats
            api_auth_accessor -> api_auth_nats

            api_user_nats -> user_service
            api_user_nats -> user_nats_server

            api_core_nats -> core_service
            api_core_nats -> core_nats_server

            api_auth_nats -> auth_service
            api_auth_nats -> auth_nats_server
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
            user -> api_service "Пользователь отправляет GET-запрос для поиска файлов по названию"
            api_service -> api_core_nats "В очередь сообщений отправляется запрос, который содержит паттерн и id пользователя, полученный из jwt токена"
            api_core_nats -> core_service "Брокер сообщений доставляет запрос в core_service"
            core_service -> api_core_nats "core_service отправляет в очередь сообщений список из ссылок на найденные файлы, составленных по их id"
            api_core_nats -> api_service "Брокер сообщений доставляет ответ от core_service"
            api_service -> user "Сервис возвращает пользователю список ссылок, по которым можно скачать файлы"
        }
    }
}
