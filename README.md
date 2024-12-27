# Calc Service

## Описание

Сервис для вычисления арифметических выражений, переданных через HTTP POST-запрос.

## Запуск

go run ./cmd/project/...

## Успешный запрос

```
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{
    "expression":"2*(3/4)+(1-3+(4/2*3))*5"
}'
```
Ответ:
```
{"result":"21.5"}
```

## Ошибка 422 (некорректное выражение)

```
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{
    "expression":"5+5*2/(5-5)"
}'
```
Ответ:
```
{"error":"Expression is not valid. Division by zero."}
```
```
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{
    "expression":"7+(2+y)"
}'
```
Ответ:
```
{"error":"Expression is not valid. Only numbers and arithmetic operations are allowed."}
```
```
 curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{
    "expression":"7++7"
}'
```
Ответ:
```
{"error":"Expression is not valid. Not enough values."}
```
## Ошибка 500 (внутренняя ошибка)

```
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{
    "expression":""
}'
```
Ответ:
```
{"error":"Internal server error"}
```
