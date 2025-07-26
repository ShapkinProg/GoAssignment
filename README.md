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
Собранный Frontend находится в соотвествующей директории.
Для пересбора:
```
cd frontend
```
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
