# Course Tasks Implementation

Этот репозиторий содержит примеры идеальной архитектуры для 10 задач курса.

## Структура (Clean Architecture)
- `api/proto/` - Контракты gRPC. Отсюда генерируется код.
- `cmd/` - Точки входа в приложение (main.go).
- `internal/domain/` - Бизнес-сущности (Car, User, Document).
- `internal/repository/` - Работа с БД (PostgreSQL, sqlc) и транзакциями (Задача 3, 8).
- `internal/service/` - Бизнес-логика, Circuit Breaker (Задача 7), Background Workers (Задача 4).
- `internal/handler/grpc/` - Обработчики gRPC запросов, Стриминг (Задача 5, 9), Interceptor (Задача 5).

## Задачи реализованы в:
- **Task 1:** `internal/repository/garage_memory.go` (In-memory map)
- **Task 3:** `internal/repository/factory_postgres.go` (Транзакции PostgreSQL)
- **Task 4:** `internal/service/warehouse_worker.go` (Фоновые воркеры)
- **Task 5:** `internal/handler/grpc/interceptor.go` (Auth Interceptor)
- **Task 8:** `internal/repository/booking_postgres.go` (Pessimistic Locking / FOR UPDATE)
- **Task 9:** `internal/handler/grpc/telemetry.go` (Bidirectional stream)
