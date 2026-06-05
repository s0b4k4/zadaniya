# Полная реализация 10 задач курса

Этот репозиторий содержит комплексное решение всех 10 задач. Проект построен по принципам Clean Architecture и использует gRPC + REST (grpc-gateway).

## Структура проекта
- `api/proto/` - Protobuf контракты для всех сервисов (Автомобили, Склад, Умный дом и т.д.)
- `cmd/main.go` - Точка входа. Мультиплексирует gRPC на порту `50051` и REST-шлюз на порту `8080`.
- `internal/domain/` - Бизнес-сущности.
- `internal/repository/` - Слои работы с БД (In-memory, PostgreSQL).
- `internal/service/` - Бизнес-логика, Background Workers (Task 4), Circuit Breaker (Task 7).
- `internal/handler/grpc/` - gRPC-обработчики, Interceptor (Task 5), Streaming (Task 9).
- `migrations/schema.sql` - Полная структура БД для всех задач.
- `sqlc/query.sql` - Запросы, из которых генерируется типобезопасный Go код.
- `docker-compose.yml` - Инфраструктура (PostgreSQL, RabbitMQ).
- `generate.bat` - Скрипт для автоматической генерации кода gRPC и sqlc.

## Как запустить на другом ПК (Инструкция)

### Шаг 1: Установка инструментов
Убедитесь, что на ПК установлены:
1. **Go** (версия 1.22+)
2. **Protoc** (компилятор Protobuf)
3. **Docker** (для запуска PostgreSQL и RabbitMQ)

### Шаг 2: Генерация кода
Просто запустите скрипт генерации (он сам скачает плагины `protoc-gen-go`, `protoc-gen-go-grpc`, `grpc-gateway`, `sqlc` и сгенерирует весь код из `.proto` и `.sql` файлов):
```cmd
generate.bat
```

### Шаг 3: Запуск инфраструктуры (БД и Очереди)
Поднимаем PostgreSQL и RabbitMQ:
```cmd
docker-compose up -d
```
(При старте PostgreSQL автоматически выполнит скрипт из `migrations/schema.sql` и создаст все таблицы).

### Шаг 4: Скачивание зависимостей
```cmd
go mod tidy
```

### Шаг 5: Запуск сервиса
```cmd
go run cmd/main.go
```

Приложение запустится:
- **gRPC сервер**: `localhost:50051`
- **REST Gateway**: `localhost:8080` (HTTP POST/GET и т.д. транслируются в gRPC)

## Соответствие решения задачам:
- **Задача 1:** `api/proto/car.proto`, `internal/repository/garage_memory.go` (in-memory)
- **Задача 2:** `api/proto/document.proto`, `migrations/schema.sql` (users, documents)
- **Задача 3:** `api/proto/car_services.proto` (ConveyorService), `sqlc/query.sql` (AssembleCar, GetCarSpec с JOIN)
- **Задача 4:** `internal/service/warehouse_worker.go` (Фоновый worker), `api/proto/warehouse.proto`
- **Задача 5:** `internal/handler/grpc/interceptor.go` (Auth), `api/proto/smarthome.proto` (Server Stream)
- **Задача 6:** `sqlc/query.sql` (`GetOwnerDashboard` — агрегированный ответ CQRS), `api/proto/car_services.proto` (DashboardService)
- **Задача 7:** `internal/service/search_client.go` (gobreaker Circuit Breaker)
- **Задача 8:** `sqlc/query.sql` (`CheckCarAvailability` содержит `FOR UPDATE` для блокировки БД при бронировании)
- **Задача 9:** `api/proto/telemetry.proto` (Bidirectional Stream `TelemetryStream`), `internal/handler/grpc/telemetry.go`
- **Задача 10:** `migrations/schema.sql` (`changelogs` таблица для логов аудита), `api/proto/car_services.proto` (Экспорт)
