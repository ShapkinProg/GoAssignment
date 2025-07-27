Веб-приложения на **Go**, реализация тестового задания.

## Запуск в Docker
В корне проекта выполнить:

```bash
docker-compose up --build
```
Сервер будет доступен по адресу:
```
http://localhost:8080 
```
## Запуск фронта:
Перед первым запуском:
```
cd frontend
npm install
```
Запуск в режиме разработки, настроен прокси на localhost:8080
```
npm run dev
```
Сборка проекта
```
npm run build
```
## Локальный запуск без Docker

1. **Создать базу данных вручную:**

```sql
CREATE DATABASE go_demo;
```
Убедитесь, что параметры подключения соответствуют:
- user:     postgres
- password: 1111
- host:     localhost
- port:     5432

2. **Запуск:**
```
go run main.go --local
```
или
```
go run main.go -l
```
