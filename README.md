Веб-приложения на **Go**, реализация тестового задания.

## Запуск в Docker
В корне проекта выполнить:

```bash
docker-compose up --build
```
Сервер будет доступен по адресу с бека:
```
http://localhost:8080
```
Или с фронта:
```
http://localhost:5173
```
## Локальный запуск без Docker

**Запуск фронта**

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

**Запуск бекенда**

```
go run main.go --local
```
или
```
go run main.go -l
```