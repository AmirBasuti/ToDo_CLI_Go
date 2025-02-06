# ToDo

![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![License](https://img.shields.io/badge/license-MIT-blue)

## Table of Contents

- [Project Description](#project-description)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Usage](#usage)
  - [Add a new task](#add-a-new-task)
  - [Mark a task as complete](#mark-a-task-as-complete)
  - [Delete a task](#delete-a-task)
  - [List all tasks](#list-all-tasks)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)
- [Contribution Guidelines](#contribution-guidelines)
- [License](#license)

## Project Description

ToDo is a command-line interface (CLI) application built with Cobra that allows you to manage your to-do list efficiently. You can add, complete, delete, and list tasks.

## Features

- Add a new task
- Mark a task as complete
- Delete a task
- List all tasks

## Prerequisites

- Go 1.22.5 or later
- PostgreSQL database

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/AmirBasuti/ToDo_CLI_Go.git
   cd ToDo_CLI_Go
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Getting Started

1. Set up the database connection:
   - Create a `.env` file in the root directory of the project.
   - Add the following environment variables to the `.env` file:
     ```env
     DB_HOST=your_database_host
     DB_USER=your_database_user
     DB_PASSWORD=your_database_password
     DB_NAME=your_database_name
     DB_PORT=your_database_port
     ```

2. Run the application:
   ```sh
   go run main.go
   ```

## Usage

### Add a new task

```sh
ToDo add "Buy groceries"
```

### Mark a task as complete

```sh
ToDo complete 42
```

### Delete a task

```sh
ToDo delete 123
```

### List all tasks

```sh
ToDo list
```

## Configuration

The application uses a `.env` file to configure the database connection. Make sure to create this file in the root directory of the project and add the following environment variables:

```env
DB_HOST=your_database_host
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
DB_PORT=your_database_port
```

## Troubleshooting

### Common Issues

1. **Database connection error**:
   - Ensure that the database server is running and accessible.
   - Verify the database connection details in the `.env` file.

2. **Command not found**:
   - Make sure you have installed the dependencies using `go mod tidy`.
   - Ensure that you are running the commands from the project root directory.

## Contribution Guidelines

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
