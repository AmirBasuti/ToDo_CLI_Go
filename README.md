# ToDo CLI in Go

A compact task manager built with Go, Cobra, GORM, and PostgreSQL. It separates command handling from task operations and stores each task's text, completion state, and creation time in a database.

## Commands

| Command | Behavior |
|---|---|
| `add "task description"` | Create a task and print its ID. |
| `list` | Print the stored tasks in a terminal table. |
| `complete <task_id>` | Mark a task complete. |
| `delete <task_id>` | Ask for confirmation, then delete a task. |

The root command also contains an interactive menu. The explicit subcommands are the clearer interface for the current implementation.

## Requirements

- Go 1.22.5 or later, as declared in `go.mod`.
- PostgreSQL and a dedicated development database.
- A `.env` file in the directory from which the application is run.

## Build

```sh
git clone https://github.com/AmirBasuti/ToDo_CLI_Go.git
cd ToDo_CLI_Go
go mod download
go build -mod=readonly ./...
```

The source compiled with Go 1.22.5 on Windows/amd64 during the 22 October 2026 review. Database commands were not executed during that review; compilation alone does not verify the database setup or command behavior.

## Database configuration

Create a local `.env` file using your own PostgreSQL connection details:

```dotenv
DB_HOST=localhost
DB_USER=<your-local-database-user>
DB_PASSWORD=<your-local-database-password>
DB_NAME=<your-development-database>
DB_PORT=5432
```

Keep this file private. The code explicitly loads `.env`, so setting environment variables alone does not replace the file requirement in this revision.

**Schema setup is currently incomplete:** `internal/Database/DataBase.go` defines a migration helper, but `main.go` does not invoke it. A fresh database therefore needs its `todos` table initialized to match `internal/Model/TodoModel.go` before task operations can work. The repository does not yet provide a verified, portable bootstrap command.

After the database and schema are prepared, the source entry point accepts these commands:

```sh
go run . add "Read the Go documentation"
go run . list
go run . complete 1
go run . delete 1
```

Use the ID returned by `add`; the example ID is illustrative. Deletion prompts for confirmation.

## Project layout

```text
main.go                    Database connection and CLI entry point
cmd/                       Cobra commands and interactive menu
internal/Database/         PostgreSQL connection and migration helper
internal/Model/            Task model
internal/TodoBehavior/     Task operations and terminal formatting
```

## Current limitations

- There are no automated test files in the current revision.
- The workflow's Go version is older than the version declared in `go.mod`; its result has not been verified here.
- Database initialization and connection-error handling need improvement before the project offers a reliable fresh-install experience.
- Interactive menu dispatch and multi-word input need validation. Quoted task descriptions in explicit subcommands follow a different input path.
- A tracked Windows executable is present, but its relationship to the current source has not been verified. Build from source when evaluating the project.

## License

The tracked `LICENSE` file is empty. No usable license grant is specified at this revision.
