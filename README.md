# Calc Service

## Описание

Сервис для вычисления арифметических выражений, переданных через HTTP POST-запрос.

## Запуск

go run ./cmd/calc_service/...

## Успешный запрос

curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "2+2*2"}'

## Ошибка 422 (некорректное выражение)

curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "2+abc"}'

## Ошибка 500 (внутренняя ошибка)

curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": ""}'
