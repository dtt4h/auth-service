# Сервис авторизации

Сервис авторизации на Go предоставляет REST API для управления пользователями, аутентификации и авторизации с использованием JWT токенов. Поддерживает регистрацию, логин, обновление токенов и валидацию.

## Описание
| Endpoint       | Method | Описание                       | Request                  | Response                                  | Auth |
| -------------- | ------ | ------------------------------ | ------------------------ | ----------------------------------------- | ---- |
| /auth/register | POST   | Регистрация пользователя       | {email, password, role?} | {access_token, refresh_token, expires_in} | No   |
| /auth/login    | POST   | Аутентификация                 | {email, password}        | {access_token, refresh_token, user_id}    | No   |
| /auth/refresh  | POST   | Обновление access_token        | {refresh_token}          | {access_token, expires_in}                | No   |
| /auth/logout   | POST   | Добавление refresh в blacklist | {refresh_token}          | 204 No Content                            | No   |
| /auth/validate | GET    | Валидация JWT + user info      | -                        | {user_id, email, roles, valid}            | JWT  |
| /profile       | GET    | Получить профиль               | -                        | {id, email, role, created_at}             | JWT  |
| /profile       | PUT    | Обновить профиль               | {email?, role?}          | {id, email, role}                         | JWT  |
| /health        | GET    | Healthcheck                    | -                        | {status: "ok", version}                   | No   |

## Структура проекта
go-auth-service/
├── cmd/
│   └── server/
│       └── main.go           # Entry point
├── internal/
│   ├── config/              # YAML/ENV config
│   ├── handler/             # HTTP handlers (Gin)
│   ├── service/             # Business logic
│   ├── repository/          # Data access (GORM/Redis)
│   ├── entity/              # DB models
│   ├── middleware/          # JWT, logging, recovery
│   └── lib/                 # Utilities (logger, validator)
├── pkg/                     # Reusable (token, crypto)
├── migrations/              # DB schema
├── tests/                   # Integration tests
├── docker-compose.yml       # Local stack
├── Dockerfile
├── Makefile
└── README.md

## Технический стэк
| Компонент | Технология              | Версия | Почему                   |
| --------- | ----------------------- | ------ | ------------------------ |
| Language  | Go                      | 1.22+  | Performance, concurrency |
| Framework | Gin                     | v1.9+  | Lightweight, middleware  |
| ORM       | GORM                    | v1.25+ | PostgreSQL migrations    |
| Cache     | go-redis                | v9.4+  | Blacklist tokens         |
| JWT       | golang-jwt              | v5.2+  | RS256 + claims           |
| Logger    | go.uber.org/zap         | v1.26+ | Structured logging       |
| Validator | go-playground/validator | v10+   | Request validation       |
| Keycloak  | oidc-client             | -      | OIDC integration         |
| DB        | PostgreSQL 16           | -      | ACID, JSONB              |
| Container | Docker Compose          | v2.27+ | Local development        |

## Предварительные требования

- Go 1.22+
- Docker & Docker Compose v2.27+
- PostgreSQL 16
- Redis (для blacklist токенов)

## Установка

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/your-repo/go-auth-service.git
   cd go-auth-service
   ```

2. Установите зависимости:
   ```bash
   go mod download
   ```

3. Запустите локальный стек:
   ```bash
   docker-compose up -d
   ```

4. Запустите миграции:
   ```bash
   make migrate-up
   ```

5. Запустите сервер:
   ```bash
   make run
   ```

## Конфигурация

Сервис использует переменные окружения для конфигурации. Создайте файл `.env` в корне проекта:

```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=auth_user
DB_PASSWORD=auth_pass
DB_NAME=auth_db

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT
JWT_SECRET=your-secret-key
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=7d

# Server
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# Logging
LOG_LEVEL=info
```

## Использование

### Регистрация пользователя
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123", "role": "user"}'
```

### Аутентификация
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123"}'
```

### Валидация токена
```bash
curl -X GET http://localhost:8080/auth/validate \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

### Получение профиля
```bash
curl -X GET http://localhost:8080/profile \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

### Обновление профиля
```bash
curl -X PUT http://localhost:8080/profile \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -d '{"email": "newemail@example.com", "role": "admin"}'
```

### Healthcheck
```bash
curl -X GET http://localhost:8080/health
```

## Тестирование

### Модульные тесты
```bash
make test-unit
```

### Интеграционные тесты
```bash
make test-integration
```

### Все тесты
```bash
make test
```

## Вклад

1. Форкните проект
2. Создайте ветку для вашей фичи (`git checkout -b feature/AmazingFeature`)
3. Зафиксируйте изменения (`git commit -m 'Add some AmazingFeature'`)
4. Запушьте ветку (`git push origin feature/AmazingFeature`)
5. Откройте Pull Request

## Список заданий

[ ] 1.1 Инициализация проекта
[ ] 1.2 Зависимости (go.mod)
[ ] 1.3 Слой конфигурации (internal/config)
[ ] 1.4 Сущности (internal/entity)
[ ] 1.5 Цели Makefile
[ ] 2.1 Настройка PostgreSQL
[ ] 2.2 Интерфейс репозитория (internal/repository)
[ ] 2.3 Реализация на GORM
[ ] 2.4 Redis-репозиторий
[ ] 2.5 Docker Compose (локальный стек)
[ ] 3.1 Утилиты JWT (pkg/token)
[ ] 3.2 Утилиты для паролей (pkg/crypto)
[ ] 3.3 Сервис аутентификации (internal/service)
[ ] 3.4 DTO и валидация
[ ] 4.1 Middleware (internal/middleware)
[ ] 4.2 Обработчики (internal/handler)
[ ] 4.3 Настройка роутера (cmd/server/main.go)
[ ] 5.1 Модульные тесты (табличные)
[ ] 5.2 Интеграционные тесты (testcontainers-go)
[ ] 5.3 OpenAPI/Swagger
[ ] 5.4 Мониторинг (Prometheus)
[ ] 5.5 README.md и Демо
[ ] 5.6 CI/CD (GitLab CI)
[ ] 5.7 Релиз (goreleaser)