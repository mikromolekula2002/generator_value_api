# Этап сборки
FROM golang:1.21-alpine AS builder

WORKDIR /usr/local/src 
#указываем рабочую директорию

RUN apk --no-cache add bash git make gcc gettext musl-dev 
#прописываем все утилиты нужные для сборки

COPY ["go.mod", "go.sum", "./"] 
#копируем проект в докер

RUN go mod download 
#качаем зависимости

COPY . . 
#копируем все в докер образ

# Компиляция приложения
RUN go build -o ./bin/app ./cmd/main.go

# Финальный образ
FROM alpine

# Копируем бинарный файл из этапа сборки
COPY --from=builder /usr/local/src/bin/app /app

# Копируем конфигурационный файл
COPY config/config.yaml /config.yaml

# Указываем команду для запуска приложения
CMD ["/app"]