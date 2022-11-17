# USER_API


## Описание
Приложение представляет собой API по работе с сущностью User, где хранилищем выступает файл json

## Задача
Сделать рефакторинг приложения (Оригинальный файл `refactoring.go`)

## Ограничения:
- Хранилищем должен оставаться файл в json формате
- Структура пользователя не должна быть уменьшена
- Приложение не должно потерять существующую функциональность


## Что было сделано
- [x] Переосмыслена архитектура проекта [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
    1. `entities` - слой работы с сущностями
    2. `adapters` - слой работы с бд (json файлом)
    3. `controller` - слой работы с http
- [x] Добавлен интерфейс `UserStore interface`, чтобы слой controller работал с абстракцией
- [x] Добавлены сущности DTO, которые передаются на уровне `handler`
- [x] Добавлен свой обработчик ошибок для ответа клиенту
- [x] Методы `handler` декомпозированы, что улучшает читаемость кода
- [x] Переписаны методы и декомпозированы, что добавляет читаемости код
- [x] Добавил файл конфигурации и логер

## Будущее
- [] Написать тесты для каждого уровня архитектуры
- [] Написать свою обертку для логера
- [] Добавил бы композиты

## Стек
- `Go`
- Фреймворк [Chi](https://github.com/go-chi/chi)
- `Docker`
- Конфигурация приложения [cleanenv](https://github.com/ilyakaznacheev/cleanenv)
- Логер [logrus](https://github.com/sirupsen/logrus)

## Запуск
1. Склонировать репозиторий
```
git clone https://github.com/vladjong/user_balance.git
```
2. Открыть терминал и набрать:
```
make
```

## Тестирование

### `http/create_user.http`
```
POST http://localhost:3333/api/v1/users
Content-Type: application/json

{
  "display_name": "TEST14",
  "email": "test14"
}
```

### `http/delete_user.http`
```
DELETE http://localhost:3333/api/v1/users/1
```

### `http/get_user.http`
```
GET http://localhost:3333/api/v1/users/1
Accept: application/json
```

### `http/search_user.http`
```
GET http://localhost:3333/api/v1/users
Accept: application/json
```

### `http/update_user.http`
```
PATCH http://localhost:3333/api/v1/users/3
Content-Type: application/json

{
  "display_name": "TEST"
}
```