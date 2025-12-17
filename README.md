# User-DOB API

### Overview
- A simple RESTful API to manage users and compute their age from date of birth (DOB).
- Built with Go using Fiber, PostgreSQL, and sqlc. Structured with clean layers: routes → handlers → services → repositories → database.

### Tech stack
- Go 1.24+
- Fiber (HTTP framework)
- PostgreSQL (database)
- sqlc (type-safe SQL → Go)
- Zap (structured logging)
- godotenv (env loading)

### Project structure
```
cmd/server/main.go            # app entrypoint
db/migrations/                # SQL migrations (up/down)
db/queries/                   # hand-written SQL for sqlc
db/sqlc/                      # sqlc generated code
internal/handler/             # HTTP handlers
internal/logger/              # zap logger factory
internal/middleware/          # request ID + request logger
internal/models/              # request/response and domain models
internal/repository/          # DB repository using sqlc
internal/routes/              # route registration
internal/service/             # business logic (age calc, user service)
sqlc.yaml                     # sqlc config
```

### Prerequisites
- Go 1.24 or newer
- PostgreSQL 13+ running and accessible

### Configuration
Create a `.env` file in the project root (you can copy from `.env.example`).
```
APP_PORT=8080
DB_URL="postgres://USER:PASSWORD@HOST:PORT/DBNAME?sslmode=disable"
```
### Notes:
- `APP_PORT` the server listens on port 8080 by default.
- `DB_URL` must be a valid PostgreSQL connection string.

### Database setup
1) Create a database in PostgreSQL (example uses db name `userdob`).
```
createdb userdob
```
2) Apply migrations (manually or with your tool of choice). Minimal manual approach:
```
psql "$DB_URL" -f db/migrations/001_create_users.up.sql
```

### Running the server
1) Install dependencies (Go will auto-resolve).
2) Start:
```
go run ./cmd/server
```
The server listens on `:8080`.

### Logging & middleware
- Each request receives an `X-Request-ID` header.
- Requests are logged with Zap (method, path, status, duration, request id).

### API
Base URL: `http://localhost:8080`

- POST `/users/` – Create user
  - Body JSON:
    ```json
    { "name": "Alice", "dob": "1990-05-23" }
    ```
  - Rules: `name` required, min length 2; `dob` required in `YYYY-MM-DD`.
  - Responses:
    - 201 Created
      ```json
      { "id": 1, "name": "Alice", "dob": "1990-05-23" }
      ```
    - 400 Bad Request (invalid JSON or date format)
    - 422 Unprocessable Entity (validation errors)

- GET `/users/` – List users
  - Response 200 OK:
    ```json
    [
      { "id": 1, "name": "Alice", "dob": "1990-05-23", "age": 34 }
    ]
    ```

- GET `/users/{id}` – Get user by id
  - Response 200 OK:
    ```json
    { "id": 1, "name": "Alice", "dob": "1990-05-23", "age": 34 }
    ```
  - 404 Not Found if user does not exist

- PUT `/users/{id}` – Update user
  - Body JSON:
    ```json
    { "name": "Alice B.", "dob": "1991-01-01" }
    ```
  - Responses:
    - 200 OK
      ```json
      { "id": 1, "name": "Alice B.", "dob": "1991-01-01" }
      ```
    - 400 Bad Request (invalid JSON or date format)
    - 404 Not Found
    - 422 Unprocessable Entity

- DELETE `/users/{id}` – Delete user
  - Responses:
    - 204 No Content
    - 404 Not Found

### Development notes
- SQL is defined in `db/queries/*.sql` and compiled to Go via `sqlc` into `db/sqlc/`.
- If you change SQL, regenerate:
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
sqlc generate
```

