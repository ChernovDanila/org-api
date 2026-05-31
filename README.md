# Org API

REST API для управления организационной структурой (подразделения и сотрудники).

## Технологии

- Go (net/http)
- PostgreSQL
- GORM
- goose (миграции)
- Docker + docker-compose

## Запуск

```bash
docker-compose up --build
```

API будет доступен на http://localhost:8080

## Миграции

```bash
goose -dir migrations postgres "host=localhost port=5432 user=postgres password=postgres dbname=org_api sslmode=disable" up
```

## API endpoints

| Метод | URL | Описание |
|-------|-----|----------|
| POST | /departments/ | Создать подразделение |
| POST | /departments/{id}/employees/ | Создать сотрудника |
| GET | /departments/{id} | Получить подразделение |
| PATCH | /departments/{id} | Обновить подразделение |
| DELETE | /departments/{id} | Удалить подразделение |