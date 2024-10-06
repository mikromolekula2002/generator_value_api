# generator_value_api

Микросервис предназначенный для генерации значений(строковых, числовых, строково-числовых). А также получения и через GET запрос.

Для работы пользователю требуется перейти по url:

POST http://localhost:8080/api/generate/?request_id=Bonesso&type=alpha  где:

request_id=insert_field - необязательный параметр, можно указывать если хотите конкретный id
ВАЖНО: при повторном запросе по тому же request_id вы получите ТО ЖЕ САМОЕ сгенерированное значение

type=insert_field необязательный параметр, отвечает за то, из чего будет состоять сгенерированное значение
Примеры запроса: 
numeric = 654698435686
string = ASDfasdfrADFgfbdt
alphanumeric = AGVr465vt8dfb546AFS

length=insert_field - стандартное значение = 32, максимальная длина = 50


GET http://localhost:8080/api/retrieve/?request_id=insert_field  где:

request_id=insert_field - обязательный параметр без которого НЕ ПОЛУЧИТЬ сгенерированное значение


Что нужно для работы микросервиса?

по пути 
|
|- config
| |- config.yaml указать свои данные
|- docker-compose.yml указать свои данные
|- makefile указать свои данные

Внимательно читайте комментарии в этих файлах, для создания бд требуется ваши данные postgreSQL, после создания БД и проведения миграции БД нужно поменять username на postgres для работы Docker Compose


Запуск создания БД осущетвляется след. образом и применения миграций:
- пишем в терминале make setup

Запуск создания БД отдельно:
- пишем в терминале make create-db

Запуск миграций отдельно:
- пишем в терминале make migrate

Запуск Docker Compose:
- пишем в терминале make docker-run

Также можно билдить и запускать Docker образ, а после Docker Compose самому:
- пишем в терминале 
- docker build -t generator:local .
- docker compose -f docker-compose.yml up
