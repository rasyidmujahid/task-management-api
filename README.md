# Task Management API

A simple REST API for managing tasks built with Go.

## Prerequisites

- Go 1.16 or higher
- Git

## Setup

1. Install Go dependencies:
```bash
go mod tidy
```

2. Run the server:
```bash
go run main.go
```

The server will start on port 8080.

## API Endpoints

### Get All Tasks
```
GET /api/tasks
```

### Get Single Task
```
GET /api/tasks/{id}
```

### Create Task
```
POST /api/tasks
```
Request body:
```json
{
    "title": "Task title",
    "description": "Task description",
    "status": "pending"
}
```

### Update Task
```
PUT /api/tasks/{id}
```
Request body:
```json
{
    "title": "Updated title",
    "description": "Updated description",
    "status": "in_progress"
}
```

### Delete Task
```
DELETE /api/tasks/{id}
```

## Task Status Values
- pending
- in_progress
- completed

## Validation Rules
- Title: Required, max 100 characters
- Description: Optional, max 500 characters
- Status: Required, must be one of: pending, in_progress, completed 