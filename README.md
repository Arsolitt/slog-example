# Пример написания обёртки для кайфа со Slog


### Что умеет

- Пихаем всякое в контекст -> Логируем с контекстом -> Получаем в логах наше всякое 
- Меняем уровень логирования в рантайме `WithLogLevel`, в разных частях приложения он может быть разным
- В константах храним названия полей, чтобы не ошибаться
- Можем использовать `WithLogValue`, чтобы прокидывать любые значения в контекст (кроме функций, они не залогируются, но я не делал проверку)
- Можем написать хелпер `WithLog<FIELD>` для каждого поля, чтобы получить более строгую проверку типов и каждый раз не пробрасывать название поля
- Можем собрать весь контекст и залогировать его одним сообщением в самом конце
- Оборачиваем ошибки, храним контекст вместе с ошибкой -> при логировании ошибки добываем этот контекст из ошибки и получаем все данные


Пример вывода
```json
{
  "time": "2024-08-27T12:17:29.045041055+03:00",
  "level": "INFO",
  "msg": "New request",
  "request_id": "123121"
}
{
  "time": "2024-08-27T12:17:29.045077275+03:00",
  "level": "DEBUG",
  "msg": "Debug message before level changed",
  "request_id": "123121"
}
{
  "time": "2024-08-27T12:17:29.045079847+03:00",
  "level": "INFO",
  "msg": "Processing request",
  "request_id": "123121",
  "request_object": {
    "Address": {
      "Host": "localhost",
      "Port": 8080
    },
    "UserAgent": "Mozilla/5.0",
    "Path": "/home"
  }
}
{
  "time": "2024-08-27T12:17:29.045102321+03:00",
  "level": "INFO",
  "msg": "Processing user",
  "request_id": "123121",
  "request_object": {
    "Address": {
      "Host": "localhost",
      "Port": 8080
    },
    "UserAgent": "Mozilla/5.0",
    "Path": "/home"
  },
  "user_id": "42"
}
{
  "time": "2024-08-27T12:17:29.045105185+03:00",
  "level": "INFO",
  "msg": "Processing instance",
  "user_id": "42",
  "instance_id": "228",
  "request_id": "123121",
  "request_object": {
    "Address": {
      "Host": "localhost",
      "Port": 8080
    },
    "UserAgent": "Mozilla/5.0",
    "Path": "/home"
  }
}
{
  "time": "2024-08-27T12:17:29.045112907+03:00",
  "level": "ERROR",
  "msg": "another error wrapping: error wrapping: some error",
  "instance_id": "228",
  "request_id": "123121",
  "request_object": {
    "Address": {
      "Host": "localhost",
      "Port": 8080
    },
    "UserAgent": "Mozilla/5.0",
    "Path": "/home"
  },
  "user_id": "42"
}
{
  "time": "2024-08-27T12:17:29.045115601+03:00",
  "level": "INFO",
  "msg": "Done",
  "user_id": "42",
  "instance_id": "228",
  "request_id": "123121",
  "request_object": {
    "Address": {
      "Host": "localhost",
      "Port": 8080
    },
    "UserAgent": "Mozilla/5.0",
    "Path": "/home"
  }
}
```