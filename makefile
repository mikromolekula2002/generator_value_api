# Укажите переменные для вашего проекта
# В поля, где прописано "insert_field" вставьте свои данные

DB_NAME := key_generator #НЕ МЕНЯТЬ!
DB_USER := postgres  #Для создания бд вписывайте настоящий юзернейм в БД, после чего меняйте на "postgres" для работы докера
DB_PASSWORD := insert_field 
DB_HOST := localhost
DB_PORT := 5432

# Команда для создания базы данных
create-db:
	psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -c "CREATE DATABASE $(DB_NAME);"

# Команда для применения миграций
migrate:
	migrate -path ./schema -database 'postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable' up

# Команда для создания базы данных и применения миграций
setup: create-db migrate

docker-run:
	docker build -t generator:local . 
	docker compose -f docker-compose.yml up