# Axonova Backend

Axonova is a high-performance backend service built with Go (Golang). It utilizes MongoDB for document storage and integrates with Gmail's SMTP server for automated email notifications.

---

## 🛠 Project Architecture

The project follows the standard Go project layout to ensure separation of concerns and maintainability.

* **cmd/**: Entry point for the application.
* **config/**: Handles configuration and environment variable loading.
* **internal/**: Contains private application and business logic (Server, Routes).
* **pkg/**: Public library code used by this project (Database connectors, Mailer).

---

## ⚙️ Prerequisites

* **Go**: 1.20+
* **MongoDB**: v4.4+ (Local instance or MongoDB Atlas)
* **SMTP Access**: A Gmail account with 2-Step Verification enabled.

---

## 🔧 Installation & Setup

### 1. Clone the Repository
```bash
git clone https://github.com/coleYab/axonovabackend.git
cd axonovabackend
```

### 2. Configure Environment Variables

Create a `.env` file in the root directory and populate it with your credentials:

```env
# Server Configuration
PORT=8080

# Database Configuration
MONGODB_URI=mongodb://127.0.0.1:27017/
MONGODB_NAME=axonova_db

# Email Configuration (Gmail)
GMAIL=your-email@gmail.com
APP_PASSWORD=your-google-app-password

```

> **Note:** To generate an `APP_PASSWORD`, go to your Google Account settings > Security > 2-Step Verification > App Passwords.

### 3. Install Dependencies

```bash
go mod tidy

```

---

## 🚀 Building and Running

### Development Mode

To run the application directly without building:

```bash
go run cmd/axonova/main.go
```

### Building for Production

To compile the source code into a single executable binary:

```bash
# Build the binary
go build -o axonova-api cmd/axonova/main.go

# Run the binary
./axonova-api

```

---

## 📖 API Documentation

The backend initializes services in the following order:

1. **Config**: Loads `.env` settings.
2. **Database**: Establishes a connection to MongoDB.
3. **Mailer**: Initializes the Gmail SMTP client.
4. **Server**: Registers routes and starts listening on the defined `PORT`.

### Core Dependencies

| Service | Package | Description |
| --- | --- | --- |
| **Storage** | `mongodb/mongo-go-driver` | High-level MongoDB driver for Go. |
| **Mailer** | `net/smtp` | Native Go package for sending emails. |
| **Config** | `joho/godotenv` | (Optional) For loading .env files. |

---

## 🤝 Contributing

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

---
