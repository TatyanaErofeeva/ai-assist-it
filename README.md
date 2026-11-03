# Todo API

REST API для управления пользователями и задачами. Сервис хранит данные в памяти процесса: после перезапуска списки пользователей и задач пусты.

Интерактивная документация доступна в Swagger UI: [http://localhost:8080/swagger/](http://localhost:8080/swagger/).

## Описание

Сервис отдаёт JSON по префиксу `/api/v1` и покрывает два ресурса:

- **users** — учётные записи с email и именем;
- **todos** — задачи, привязанные к пользователю.

Успешный ответ с телом завёрнут в объект `SuccessResponse`: флаг `success` и полезная нагрузка в `data`. Пустой список сериализуется как `null`. Удаление возвращает `204` без тела. Ошибки разбора и отсутствующие объекты приходят как текст (`text/plain`).

## Установка и запуск

Нужны Go 1.22 или новее и CLI [swag](https://github.com/swaggo/swag).

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Убедитесь, что `$GOPATH/bin` (обычно `~/go/bin`) есть в `PATH`.

Перейдите в каталог с `go.mod` и соберите спецификацию. Обработчики и модели лежат в `internal/`, поэтому swag запускается с `--parseInternal`:

```bash
swag init -g main.go -o docs --parseInternal
```

Команда создаёт `docs/swagger.json`, `docs/swagger.yaml` и `docs/docs.go`.

Запустите сервер:

```bash
go run main.go
```

- API: [http://localhost:8080](http://localhost:8080)
- Swagger UI: [http://localhost:8080/swagger/](http://localhost:8080/swagger/)
- Спецификация: [http://localhost:8080/docs/swagger.json](http://localhost:8080/docs/swagger.json)

Порт `8080` должен быть свободен.

## Зависимости

| Зависимость | Версия | Назначение |
| --- | --- | --- |
| Go | 1.22+ | Язык и toolchain |
| [github.com/go-chi/chi/v5](https://github.com/go-chi/chi) | v5.0.10 | HTTP-маршрутизатор |
| [github.com/swaggo/swag](https://github.com/swaggo/swag) | v1.16.2 | Генерация OpenAPI из аннотаций |
| [github.com/swaggo/http-swagger/v2](https://github.com/swaggo/http-swagger) | v2.0.2 | Swagger UI |
| swag CLI | v1.16.2 | Команда `swag init` |

Косвенные модули перечислены в `go.mod` и `go.sum`.

## API

Базовый URL: `http://localhost:8080/api/v1`.

Полное описание схем, кодов ответа и кнопка **Try it out** находятся в [Swagger UI](http://localhost:8080/swagger/).

### Пользователи

| Метод | Путь | Описание |
| --- | --- | --- |
| `GET` | `/users` | Список пользователей |
| `POST` | `/users` | Создать пользователя |
| `GET` | `/users/{id}` | Получить пользователя |
| `PUT` | `/users/{id}` | Частично обновить пользователя |
| `DELETE` | `/users/{id}` | Удалить пользователя |

### Задачи

| Метод | Путь | Описание |
| --- | --- | --- |
| `GET` | `/todos` | Список задач |
| `POST` | `/todos` | Создать задачу |
| `GET` | `/todos/{id}` | Получить задачу |
| `PUT` | `/todos/{id}` | Частично обновить задачу |
| `DELETE` | `/todos/{id}` | Удалить задачу |

## Примеры запросов и ответов

### Создать пользователя

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H 'Content-Type: application/json' \
  -d '{"email":"anna@example.com","name":"Анна"}'
```

`201 Created`

```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "anna@example.com",
    "name": "Анна",
    "created_at": "2026-11-03T11:34:14+03:00"
  }
}
```

### Список пользователей

```bash
curl http://localhost:8080/api/v1/users
```

`200 OK`

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "email": "anna@example.com",
      "name": "Анна",
      "created_at": "2026-11-03T11:34:14+03:00"
    }
  ]
}
```

Пока пользователей нет, ответ выглядит так: `{"success":true,"data":null}`.

### Получить пользователя

```bash
curl http://localhost:8080/api/v1/users/1
```

`200 OK` — в `data` один объект `User`, как в ответе на создание.

### Обновить пользователя

В теле передаются только изменяемые поля.

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H 'Content-Type: application/json' \
  -d '{"name":"Анна Иванова"}'
```

`200 OK`

```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "anna@example.com",
    "name": "Анна Иванова",
    "created_at": "2026-11-03T11:34:14+03:00"
  }
}
```

### Удалить пользователя

```bash
curl -i -X DELETE http://localhost:8080/api/v1/users/1
```

`204 No Content`, тело пустое.

### Создать задачу

```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H 'Content-Type: application/json' \
  -d '{"user_id":1,"title":"Купить молоко","description":"2 литра, 3.2%","done":false}'
```

`201 Created`

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Купить молоко",
    "description": "2 литра, 3.2%",
    "done": false,
    "created_at": "2026-11-03T11:34:14+03:00",
    "updated_at": "2026-11-03T11:34:14+03:00"
  }
}
```

### Обновить задачу

```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H 'Content-Type: application/json' \
  -d '{"done":true}'
```

`200 OK` — в `data` задача с `"done": true` и новым `updated_at`.

### Удалить задачу

```bash
curl -i -X DELETE http://localhost:8080/api/v1/todos/1
```

`204 No Content`.

### Ошибки

```bash
curl -i http://localhost:8080/api/v1/todos/abc
```

`400 Bad Request`, тело: `Неверный ID`.

```bash
curl -i http://localhost:8080/api/v1/users/99
```

`404 Not Found`, тело: `Пользователь не найден`.

```bash
curl -i -X POST http://localhost:8080/api/v1/users \
  -H 'Content-Type: application/json' \
  -d '{'
```

`400 Bad Request`, тело: `Невалидный запрос`.

Повторное удаление уже удалённой задачи возвращает `404` и текст `Задача не найдена`.

## Чек-лист по структуре проекта

```text
.
├── main.go                     # маршруты, Swagger UI, метаданные API
├── go.mod
├── go.sum
├── README.md
├── LICENSE
├── docs/
│   ├── docs.go                 # код, сгенерированный swag
│   ├── swagger.json            # OpenAPI 2.0
│   └── swagger.yaml
└── internal/
    ├── handlers/
    │   ├── todo.go             # аннотации задач
    │   └── user.go             # аннотации пользователей
    └── models/
        ├── todo.go             # Todo и запросы задач
        └── user.go             # User, ответы и запросы пользователей
```

- [x] У каждого endpoint есть Swagger-аннотации: summary, описание, тег, параметры, коды ответа
- [x] У моделей и их полей есть godoc-комментарии
- [x] `docs/swagger.json` собран командой `swag init -g main.go -o docs --parseInternal`
- [x] Swagger UI открывается по адресу `/swagger/` и показывает группы `users` и `todos`
- [x] README содержит описание, установку, запуск, зависимости, API, примеры, структуру, лицензию и контакты

## Лицензия и контакты

Проект распространяется по лицензии MIT. Полный текст лежит в [LICENSE](LICENSE).

Вопросы по API: [support@example.com](mailto:support@example.com).
