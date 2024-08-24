# Пример написания обёртки для кайфа со Slog


### Что умеет

- Пихаем всякое в контекст -> Логируем с контекстом -> Получаем в логах наше всякое 
- Меняем уровень логирования в рантайме `WithLogLevel`, в разных частях приложения он может быть разным
- В константах храним названия полей, чтобы не ошибаться
- Можем использовать `WithLogValue`, чтобы прокидывать любые значения в контекст (в данном примере захардкодил строку, но при желании можно сделать пустой интерфейс и запилить логику на проверке типов)
- Можем написать хелпер `WithLog<FIELD>` для каждого поля, чтобы получить более строгую проверку типов и каждый раз не пробрасывать название поля
- Можем собрать весь контекст и залогировать его одним сообщением в самом конце
- Оборачиваем ошибки, храним контекст вместе с ошибкой -> при логировании ошибки добываем этот контекст из ошибки и получаем все данные


Пример вывода
```json
{
  "time": "2024-08-24T17:35:34.310888407+03:00",
  "level": "INFO",
  "msg": "New request",
  "request_id": "123121"
}
{
  "time": "2024-08-24T17:35:34.310924334+03:00",
  "level": "DEBUG",
  "msg": "Debug message before level changed",
  "request_id": "123121"
}
{
  "time": "2024-08-24T17:35:34.31092672+03:00",
  "level": "INFO",
  "msg": "Processing user",
  "request_id": "123121",
  "user_id": "42"
}
{
  "time": "2024-08-24T17:35:34.310928654+03:00",
  "level": "INFO",
  "msg": "Processing instance",
  "user_id": "42",
  "instance_id": "228",
  "request_id": "123121"
}
{
  "time": "2024-08-24T17:35:34.310940237+03:00",
  "level": "ERROR",
  "msg": "another error wrapping: error wrapping: some error",
  "request_id": "123121",
  "user_id": "42",
  "instance_id": "228"
}
{
  "time": "2024-08-24T17:35:34.310942282+03:00",
  "level": "INFO",
  "msg": "Done",
  "request_id": "123121",
  "user_id": "42",
  "instance_id": "228"
}
```