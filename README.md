# Example of Writing a Wrapper for Slog with Style

### Features

- Put values in context -> Log with context -> Get our values in logs
- Change log level at runtime with `WithLogLevel`, different parts of the application can have different levels
- Store field names in constants to avoid mistakes
- Use `WithLogValue` to pass any values into the context (except functions, they won't be logged, but I didn't add a check)
- Write helper `WithLog<FIELD>` for each field to get stricter type checking and avoid passing field name each time
- Collect all context and log it in a single message at the end
- Wrap errors, store context with the error -> when logging the error, extract this context and get all data
- Use colored PrettyHandler for local development `isPretty = true`

Example output
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

## Project Structure

- `logger/` - Core logging functionality
  - `errors.go` - Error wrapping with context
  - `logger.go` - Main logging functions
  - `middleware.go` - Slog middleware for context handling
  - `pretty.go` - Pretty formatter for console output
  - `types.go` - Type definitions and constants
- `main.go` - Example usage
- `Taskfile` - Build and run commands

## Usage

Build and run the example:

```bash
task run
```
