# avitoTestTask

# API-сервер для Управления Сегментацией Пользователей - Руководство

## Начало работы

Чтобы настроить API-сервер и начать использование, выполните следующие шаги:

1. Запуск:
   ```bash
   make compose-up
   ```

2. Соберите и запустите API-сервер:
   ```bash
   make
   ./app
   ```

API-сервер начнет работу на указанном порту (по умолчанию: 8080).

## API-точки

### Создание Сегмента

Создает новый сегмент с указанным идентификатором.

**Точка входа:** `POST /segment/{slug}`

#### Пример запроса

```bash
curl -X POST http://localhost:8080/segment/new-segment
```

#### Пример ответа

```json
{
    "id": 1,
    "slug": "new-segment",
    "created_at": "2023-08-31T15:31:19.111915Z"
}
```

### Удаление Сегмента

Удаляет существующий сегмент с указанным идентификатором.

**Точка входа:** `DELETE /segment/{slug}`

#### Пример запроса

```bash
curl -X DELETE http://localhost:8080/segment/new-segment
```

#### Пример ответа

```json
{
    "id": 1,
    "slug": "new-segment",
    "created_at": "2023-08-31T15:31:19.111915Z"
}
```

### Добавление Сегментов к Пользователю

Связывает несколько сегментов с определенным пользователем.

**Точка входа:** `POST /user/{id}`

#### Пример запроса

```bash
curl -X POST -H "Content-Type: application/json" -d '{"slugs": ["segment1", "segment2"]}' http://localhost:8080/user/123
```

#### Пример ответа

```json
[
  null,
  null
]
```

### Удаление Сегментов у Пользователя

Удаляет несколько сегментов у определенного пользователя.

**Точка входа:** `DELETE /user/{id}`

#### Пример запроса

```bash
curl -X DELETE -H "Content-Type: application/json" -d '{"slugs": ["segment1", "segment2"]}' http://localhost:8080/user/123
```

#### Пример ответа

```json
[
  null,
  null
]
```

### Получение Сегментов Пользователя

Получает все сегменты, связанные с определенным пользователем.

**Точка входа:** `GET /user/{id}`

#### Пример запроса

```bash
curl http://localhost:8080/user/123
```

#### Пример ответа

```json
[
  {
    "UserId": 123,
    "SegmentSlug": "segment1"
  },
  {
    "UserId": 123,
    "SegmentSlug": "segment2"
  }
]
```
