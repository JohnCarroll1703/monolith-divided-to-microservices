# Microservices Project

Этот проект построен с использованием **микросервисной архитектуры**.  
Каждый сервис отвечает за свою зону ответственности и взаимодействует с другими через API (HTTP/gRPC).

## 🚀 Стек технологий
- [Go](https://go.dev/) — язык реализации
- [gRPC](https://grpc.io/) — внутренняя коммуникация между сервисами
- [Gin](https://gin-gonic.com/) — REST API для CRUD операций
- [PostgreSQL](https://www.postgresql.org/) — основное хранилище данных
- [Docker](https://www.docker.com/) — контейнеризация сервисов
- [Migrations](https://github.com/golang-migrate/migrate) — управление схемой БД
- [Kafka](https://github.com/segmentio/kafka-go) — обмен событиями и стриминг данных
- [Stripe](https://stripe.com/) — интеграция платёжной системы

## ⚡ REST API
Примеры ручек (через Gin):
- `GET /users/:id` — получить пользователя
- `POST /users` — создать пользователя
- `PUT /users/:id` — обновить пользователя
- `DELETE /users/:id` — удалить пользователя
- `GET /health` — health-check

## 🔌 gRPC API
gRPC-контракты описаны в [sdk/proto]

## ✅ TODO
1. Добавить аутентификацию/авторизацию
2. Написать интеграционные тесты
3. Подключить CI/CD (GitHub Actions)
4. Логирование и метрики (Prometheus)

