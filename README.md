# Timeline API

A REST API for managing timeline wraps and media files built with Go and SQLite.

## Overview

Timeline API is a backend service that allows users to create and manage timeline wraps containing media files. Each wrap can contain up to 9 media files (images, videos, audio) and provides a simple way to organize and retrieve timeline content.

## Current State

The project is currently in active development with the following implemented features:

### Completed Features

**Wrap Management:**
- Create new timeline wraps with name and status
- Retrieve wrap information by UUID
- Delete wraps with cascade deletion of associated media
- Automatic UUID generation for wraps
- Timestamp tracking (created_at, updated_at)

**Media Management (Repository Layer):**
- Media model with support for images, videos, and audio files
- File metadata storage (filename, file path, file size, MIME type)
- Upload timestamp and photo taken timestamp tracking
- Automatic UUID generation for media files
- 9 file limit per wrap enforcement
- Cascade deletion when parent wrap is deleted

**Database:**
- SQLite database with foreign key constraints
- Automatic table creation and migration
- Proper indexing and relationships between wraps and media

**API Endpoints:**
- POST /api/create/wrap - Create a new wrap
- GET /api/get/wrap/{uuid} - Retrieve wrap by UUID
- GET /health - Health check endpoint
- GET / - Root endpoint with API status

**Infrastructure:**
- Deployed on Render cloud platform
- Environment-based port configuration
- Structured logging and error handling
- Clean architecture with separation of concerns (handlers, services, repositories)

### Architecture

The project follows a clean architecture pattern with the following layers:

```
cmd/api/           - Application entry point
internal/
  db/              - Database connection and migrations
  wrap/            - Wrap domain (model, repository, service, handler)
  media/           - Media domain (model, repository, service)
```

## Technology Stack

- **Language:** Go 1.21+
- **Database:** SQLite3
- **HTTP Router:** Go standard library net/http
- **UUID Generation:** google/uuid package
- **Database Driver:** mattn/go-sqlite3
- **Deployment:** Render

## Installation and Setup

### Prerequisites

- Go 1.21 or higher
- SQLite3

### Local Development

1. Clone the repository:
```bash
git clone <repository-url>
cd timeline
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run cmd/api/main.go
```

The server will start on port 8080 by default.

### Environment Variables

- `PORT` - Server port (defaults to 8080)

## API Usage

### Create a Wrap

```bash
POST /api/create/wrap
Content-Type: application/json

{
    "name": "My Timeline Wrap",
    "status": "active"
}
```

**Response:**
```json
{
    "uuid": "generated-uuid",
    "name": "My Timeline Wrap",
    "status": "active",
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
}
```

### Get a Wrap

```bash
GET /api/get/wrap/{uuid}
```

**Response:**
```json
{
    "uuid": "wrap-uuid",
    "name": "My Timeline Wrap",
    "status": "active",
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
}
```

### Health Check

```bash
GET /health
```

**Response:**
```
OK
```

## Development Status

### What's Working
- Complete wrap CRUD operations
- Database persistence with SQLite
- RESTful API endpoints
- Proper error handling and validation
- UUID-based resource identification
- Foreign key constraints and cascade deletion
- Production deployment on Render

### Current Development Focus
- Media upload and management endpoints
- File storage integration
- Media retrieval by wrap
- Enhanced error responses and validation

## Deployment

The application is currently deployed on Render at:
```
https://timeline-zuv1.onrender.com
```

### Deployment Configuration

The application automatically:
- Detects the PORT environment variable from Render
- Initializes SQLite database on startup
- Creates required tables if they don't exist
- Enables foreign key constraints for data integrity

## Testing

The API has been tested using:
- Postman for POST requests and JSON payload testing
- Browser testing for GET endpoints
- Production deployment verification on Render

### Test Endpoints

- Health check: `GET https://timeline-zuv1.onrender.com/health`
- Create wrap: `POST https://timeline-zuv1.onrender.com/api/create/wrap`
- Get wrap: `GET https://timeline-zuv1.onrender.com/api/get/wrap/{uuid}`

## Project Structure

```
timeline/
├── cmd/api/main.go                 # Application entry point
├── internal/
│   ├── db/
│   │   ├── db_connection.go        # Database connection setup
│   │   └── migrations.go           # Table creation and schema
│   ├── wrap/
│   │   ├── model.go               # Wrap data structure
│   │   ├── repository.go          # Database operations
│   │   ├── service.go             # Business logic
│   │   ├── handler.go             # HTTP handlers
│   │   └── interface.go           # Repository interface
│   └── media/
│       ├── model.go               # Media data structure
│       ├── repository.go          # Database operations
│       └── service.go             # Business logic
├── media/                         # Media file storage directories
│   ├── images/
│   ├── videos/
│   └── audio/
├── go.mod                         # Go module definition
├── go.sum                         # Dependency checksums
└── README.md                      # Project documentation
```