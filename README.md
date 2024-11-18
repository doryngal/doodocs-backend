# 🌟 Doodocs Backend API

Welcome to the **Doodocs Backend API**! This project is a robust REST API designed for managing archives and sending emails. It was developed as part of the **Doodocs Backend Challenge**. 🚀

---

## ✨ Features

- 📁 **Analyze archive files**: Extract detailed structure and metadata.
- 🗜️ **Create ZIP archives**: Combine multiple files into a single archive.
- 📧 **Send files via email**: Deliver files to multiple recipients using SMTP.

---

## 🛠️ Technologies Used

- 💻 **Go (Golang)**: Backend logic.
- 🌐 **Gin**: HTTP web framework.
- 📖 **Swaggo**: API documentation generation.
- ✉️ **SMTP**: Email functionality.
- 🐳 **Docker**: Containerization.

---

## 📂 Project Structure

```plaintext
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
│   ├── testdata/        # Sample files for testing
├── Dockerfile           # Dockerfile for containerization
├── go.mod               # Go module dependencies
├── go.sum               # Go dependency checksum file
└── README.md            # Project documentation
```
## ⚙️ Setup Instructions

### 📋 Prerequisites

- 🛠️ **Go 1.22 or higher**
- 🐳 **Docker** (optional for containerization)

---

### 🚀 Local Setup

1. **Clone the repository**:
    ```bash
    git clone https://github.com/doryngal/doodocs-backend.git
    cd doodocs-backend
    ```

2. **Copy the `.env` file and update values**:
    ```bash
    cp .env.example .env
    ```

3. **Install dependencies**:
    ```bash
    go mod download
    ```

4. **Run the application**:
    ```bash
    go run cmd/main.go
    ```

5. **Access Swagger Documentation**:
   Open [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) in your browser.

---

### 🐳 Using Docker

1. **Build the Docker image**:
    ```bash
    docker build -t doodocs-backend .
    ```

2. **Run the Docker container**:
    ```bash
    docker run -p 8080:8080 --env-file .env doodocs-backend
    ```

---

### 🧪 Testing

#### **Unit Tests**

1. **Run tests**:
    ```bash
    go test ./...
    ```

2. **What’s tested?**:
    - 📊 **Archive Analysis Validation**: Ensure archive files are analyzed correctly.
    - 🗂️ **Archive Creation**: Validate handling of valid and invalid files during archive creation.
    - ✉️ **Email Functionality**: Test email-sending with mocked configurations.

---

### 🛡️ Guidelines Followed

- 💻 **Clean Code**: Modular and maintainable design.
- 🏛️ **Clean Architecture**: Decoupled service and controller logic.
- 🧩 **Extensibility**: Easy to add new features.
- 🧪 **Testing**: Comprehensive unit tests for all key functionalities.
- 🚀 **Performance**: Optimized handling of archives and email processes.
