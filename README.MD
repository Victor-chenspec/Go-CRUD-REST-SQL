# Task API (Go)

REST API for users and their tasks. Built with [Gin](https://github.com/gin-gonic/gin) and PostgreSQL via [pgx](https://github.com/jackc/pgx).

## Layout

```
cmd/server/          # application entrypoint
internal/handler/    # HTTP handlers
internal/service/    # business layer
internal/repository/ # PostgreSQL queries
internal/model/      # request/response structs
apperror/            # typed HTTP errors
middleware/          # Gin error handler
```

User handlers attach errors with `ctx.Error`. A global `ErrorHandler` middleware then writes the JSON response. Task handlers still write status and body directly; the payload shape is the same (`{"error": "..."}`).

## Prerequisites

- Go (see `go.mod`)
- PostgreSQL
- A `.env` file in the project root — `godotenv.Load()` exits the process if the file cannot be loaded

## Environment

Create `.env`:

```env
PORT=:8080
SQL=postgres://USER:PASSWORD@localhost:5432/DATABASE?sslmode=disable
```

| Variable | Description |
|----------|-------------|
| `PORT` | Address passed to Gin, e.g. `:8080` |
| `SQL` | PostgreSQL connection string for pgx |

## Database schema

The app expects these tables. Create them before starting the server. Column order matches the repository `Scan` calls (`SELECT *`).

```sql
CREATE TABLE users (
    id    SERIAL PRIMARY KEY,
    name  TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);

CREATE TABLE tasks (
    id        SERIAL PRIMARY KEY,
    title     TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    user_id   INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE
);
```

## Run

```bash
go mod download
go run ./cmd/server
```

The process loads `.env`, connects to Postgres, registers `middleware.ErrorHandler()`, then starts Gin on `PORT`.

## API

Base URL: `http://localhost:8080` (if `PORT=:8080`).

### Users

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/user` | List all users |
| `GET` | `/user/:id` | Get one user |
| `POST` | `/user` | Create user |
| `PUT` | `/user` | Update user (`id` in JSON body) |
| `DELETE` | `/user/:id` | Delete user |

**Create user**

```http
POST /user
Content-Type: application/json

{
  "name": "Ada",
  "email": "ada@example.com"
}
```

**Update user**

```http
PUT /user
Content-Type: application/json

{
  "id": 1,
  "name": "Ada Lovelace",
  "email": "ada@example.com"
}
```

**Response shape**

```json
{
  "id": 1,
  "name": "Ada",
  "email": "ada@example.com"
}
```

### Tasks

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/task/:id` | List tasks for **user** `:id` |
| `POST` | `/task` | Create task |
| `PUT` | `/task` | Update task (`id` in JSON body) |
| `DELETE` | `/task/:id` | Delete task by **task** id |

`GET /task/:id` uses the path id as `user_id`. `DELETE /task/:id` uses the path id as the task's primary key.

**Create task**

```http
POST /task
Content-Type: application/json

{
  "title": "Write README",
  "user_id": 1
}
```

New tasks use the database default for `completed` (typically `false`).

**Update task**

```http
PUT /task
Content-Type: application/json

{
  "id": 1,
  "title": "Write README",
  "completed": true,
  "user_id": 1
}
```

**Response shape**

```json
{
  "id": 1,
  "title": "Write README",
  "completed": false,
  "user_id": 1
}
```

### Errors

Failure bodies look like:

```json
{
  "error": "User not found"
}
```

| Code | When |
|------|------|
| `200` | Success |
| `201` | Created |
| `404` | User/task not found (user routes: `apperror.NotFound`; task update/delete still return `"User not found"`) |
| `422` | Invalid id or JSON / validation |
| `500` | Database error, or an untyped error caught by middleware (`"Internal server error"`) |

User service maps repository failures to `apperror.InternalError("Database error")` and missing rows to `apperror.NotFound("User not found")`. Task routes do not use `apperror` yet.

## Example curl

```bash
# Create user
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Ada","email":"ada@example.com"}'

# List users
curl http://localhost:8080/user

# Get one user
curl http://localhost:8080/user/1

# Update user
curl -X PUT http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"id":1,"name":"Ada Lovelace","email":"ada@example.com"}'

# Create task
curl -X POST http://localhost:8080/task \
  -H "Content-Type: application/json" \
  -d '{"title":"Write README","user_id":1}'

# List tasks for user 1
curl http://localhost:8080/task/1

# Update task
curl -X PUT http://localhost:8080/task \
  -H "Content-Type: application/json" \
  -d '{"id":1,"title":"Write README","completed":true,"user_id":1}'

# Delete task
curl -X DELETE http://localhost:8080/task/1

# Delete user
curl -X DELETE http://localhost:8080/user/1
```

## License

MIT — see [LICENSE](LICENSE).
