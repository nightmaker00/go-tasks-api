# go-tasks-api

go-tasks-api — REST API сервис для управления задачами (CRUDL),
написанный на чистом Go с использованием стандартной библиотеки.

Проект создан как pet-проект для практики:
- net/http
- работы с PostgreSQL
- Docker / docker-compose
- базовой backend-архитектуры

## Стек технологий
- Go 1.22
- net/http (без фреймворков)
- PostgreSQL
- database/sql
- Docker, Docker Compose

Используется только стандартная библиотека Go (кроме драйвера PostgreSQL).

## Быстрый старт

1) Создать `.env` из шаблона:

```
cp .env.example .env
```