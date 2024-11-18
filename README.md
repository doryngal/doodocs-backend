Doodocs Backend API

This project is a REST API for handling archives and sending emails, developed as part of the Doodocs Backend Challenge.

Features

	1.	Analyze archive files to extract detailed structure and metadata.
	2.	Create and return a ZIP archive from uploaded files.
	3.	Send files to multiple email addresses via SMTP.

Technologies Used

	•	Go (Golang): For backend logic.
	•	Gin: HTTP web framework.
	•	Swaggo: For API documentation generation.
	•	SMTP: For email functionality.
	•	Docker: For containerizing the application.

Project Structure
```
doodocs-backend/
├── cmd/                 # Main application entry point
├── internal/
│   ├── config/          # Configuration and environment loading
│   ├── controller/      # Route handlers
│   ├── docs/            # Swagger documentation (auto-generated)
│   ├── middleware/      # Middleware logic
│   ├── model/           # Data models
│   ├── service/         # Business logic
├── test/                # Tests
│   ├──testdata/         # Sample files for testing
├── Dockerfile           # Dockerfile for containerization
├── go.mod               # Go module dependencies
├── go.sum               # Go dependency checksum file
└── README.md            # Project documentation
```

Setup Instructions

Prerequisites

    •	Go 1.22 or higher.
    •	Docker (optional for containerization).

Local Setup

1.	Clone the repository:

```bash
git clone https://github.com/your-repo/doodocs-backend.git
cd doodocs-backend
```

2.	Copy the example .env file and update the values:
```bash
cp .env.example .env
```

3.	Install dependencies:
```bash
go mod download
```

4.	Run the application:
```bash
go run cmd/main.go
```

5.	Access Swagger documentation:
```
http://localhost:8080/swagger/index.html
```
Using Docker

1.	Build the Docker image:
```bash
docker build -t doodocs-backend .
```

2.	Run the Docker container:
```bash
docker run -p 8080:8080 --env-file .env doodocs-backend
```
Testing

Unit Tests

1.	Run unit tests:
```bash
go test ./...
```

2.	Tests include:

	•	Validating archive analysis.
	•	Ensuring archive creation handles valid and invalid input.
	•	Testing email functionality with mocked configurations.

Guidelines Followed

    •	Clean Code: Modular design with clear separation of concerns.
	•	Clean Architecture: Business logic is isolated in service layers.
	•	Extensibility: Easy to add new features.
	•	Tests: Comprehensive unit tests for core functionality.
	•	Optimized: Efficient handling of archives and email processes.

