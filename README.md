# Wallet Service

[![Go Report Card](https://goreportcard.com/badge/github.com/your-username/wallet-service)](https://goreportcard.com/report/github.com/your-username/wallet-service)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://img.shields.io/badge/docker-ready-blue.svg)](https://www.docker.com/)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=flat&logo=postgresql&logoColor=white)](https://www.postgresql.org)

A **scalable, production-ready wallet and transaction management microservice** built with **Go, Gin, and GORM**. This service provides a robust foundation for financial applications requiring secure wallet management, transaction processing, and audit trails with strong data consistency and modular architecture.

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Prerequisites](#-prerequisites)
- [Installation](#-installation)
- [Configuration](#-configuration)
- [Usage](#-usage)
- [API Documentation](#-api-documentation)
- [Project Structure](#-project-structure)
- [Testing](#-testing)
- [Deployment](#-deployment)
- [Architecture](#-architecture)
- [Contributing](#-contributing)
- [License](#-license)
- [Support](#-support)

## ✨ Features

### Core Functionality
- **🔐 User Authentication** - JWT-based authentication with pluggable OAuth2 support
- **💳 Wallet Management** - Create, fund, withdraw, and transfer wallets with multi-currency support
- **📊 Transaction Tracking** - Complete transaction history with status tracking and audit trails
- **🔄 Real-time Balance Updates** - Atomic balance updates with strong consistency guarantees

### Technical Features
- **🗄️ Database Integration** - PostgreSQL with GORM ORM for efficient data operations
- **🌐 RESTful API** - Clean HTTP endpoints built with Gin framework
- **✅ Error Handling & Validation** - Consistent JSON error responses with comprehensive input validation
- **🏗️ Modular Architecture** - Clean separation of concerns with services, repositories, and models
- **📈 Scalability** - Designed for horizontal scaling with containerization support
- **🔍 Observability** - Structured logging and metrics ready for monitoring integration

## 🛠️ Tech Stack

| Component | Technology | Version |
|-----------|------------|---------|
| **Language** | Go (Golang) | 1.21+ |
| **Framework** | Gin | ^1.9 |
| **ORM** | GORM | ^1.25 |
| **Database** | PostgreSQL | 15+ |
| **Authentication** | JWT | - |
| **Configuration** | Viper | ^1.16 |
| **Validation** | Validator | ^10.14 |
| **Testing** | Testify | ^1.8 |
| **Containerization** | Docker | 20+ |

## ⚙️ Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.21 or higher** - [Download Go](https://golang.org/dl/)
- **PostgreSQL 15 or higher** - [Download PostgreSQL](https://www.postgresql.org/download/)
- **Docker & Docker Compose** (optional) - [Download Docker](https://www.docker.com/products/docker-desktop)
- **Git** - [Download Git](https://git-scm.com/downloads)

## 🚀 Installation

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/wallet-service.git
cd wallet-service
```

### 2. Install Dependencies

```bash
# Install Go dependencies
go mod tidy
```

### 3. Database Setup

#### Option A: Local PostgreSQL

```bash
# Create database
createdb wallet_service

# Run migrations (if available)
go run cmd/wallet-service/main.go migrate
```

#### Option B: Docker Compose

```bash
# Start PostgreSQL with Docker Compose
docker-compose up -d postgres

# Wait for PostgreSQL to be ready
docker-compose exec postgres pg_isready
```

## ⚙️ Configuration

### Environment Variables

Create a `.env` file in the project root:

```env
# Application Configuration
APP_PORT=8080
APP_ENV=development
APP_DEBUG=true

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_secure_password
DB_NAME=wallet_service
DB_SSL_MODE=disable
DB_MAX_CONNECTIONS=25
DB_MAX_IDLE_CONNECTIONS=5
DB_CONNECTION_LIFETIME=5m

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_here
JWT_EXPIRATION=24h
JWT_REFRESH_EXPIRATION=168h

# Rate Limiting
RATE_LIMIT_REQUESTS_PER_MINUTE=60

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
```

### Configuration File

Alternatively, you can use a `config.yaml` file:

```yaml
app:
  port: 8080
  env: development
  debug: true

database:
  host: localhost
  port: 5432
  user: postgres
  password: your_secure_password
  name: wallet_service
  ssl_mode: disable
  max_connections: 25
  max_idle_connections: 5
  connection_lifetime: 5m

jwt:
  secret: your_super_secret_jwt_key_here
  expiration: 24h
  refresh_expiration: 168h

rate_limit:
  requests_per_minute: 60

logging:
  level: info
  format: json
```

## 🎯 Usage

### Running the Application

#### Development Mode

```bash
# Run with hot reload (requires air)
air -c .air.toml

# Or run directly
go run cmd/wallet-service/main.go
```

#### Production Mode

```bash
# Build the application
go build -o bin/wallet-service cmd/wallet-service/main.go

# Run the binary
./bin/wallet-service
```

#### Docker

```bash
# Build and run with Docker Compose
docker-compose up --build

# Or build Docker image manually
docker build -t wallet-service .
docker run -p 8080:8080 --env-file .env wallet-service
```

### Health Check

Verify the service is running:

```bash
curl http://localhost:8080/health
```

Expected response:

```json
{
  "status": "ok",
  "timestamp": "2024-01-01T00:00:00Z",
  "version": "1.0.0",
  "database": "connected"
}
```

## 📡 API Documentation

### Authentication

All API endpoints (except health check) require JWT authentication. Include the token in the Authorization header:

```http
Authorization: Bearer <your-jwt-token>
```

### Base URL

- **Development**: `http://localhost:8080`
- **Production**: `https://api.yourdomain.com`

### Endpoints

#### 1. Health Check

```http
GET /health
```

**Response:**
```json
{
  "status": "ok",
  "timestamp": "2024-01-01T00:00:00Z",
  "version": "1.0.0"
}
```

#### 2. User Authentication

##### Login
```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2024-01-02T00:00:00Z",
  "user": {
    "id": "123",
    "email": "user@example.com",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 3. Wallet Management

##### Create Wallet
```http
POST /api/wallets
Content-Type: application/json
Authorization: Bearer <token>

{
  "currency": "USD",
  "wallet_type": "personal"
}
```

**Response:**
```json
{
  "id": "wallet_123",
  "user_id": "user_123",
  "currency": "USD",
  "balance": 0.00,
  "status": "active",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

##### Get Wallet
```http
GET /api/wallets/{wallet_id}
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": "wallet_123",
  "user_id": "user_123",
  "currency": "USD",
  "balance": 150.50,
  "status": "active",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

##### Fund Wallet
```http
POST /api/wallets/{wallet_id}/fund
Content-Type: application/json
Authorization: Bearer <token>

{
  "amount": 100.50,
  "reference": "DEPOSIT_123",
  "description": "Bank deposit"
}
```

**Response:**
```json
{
  "transaction_id": "txn_123",
  "wallet_id": "wallet_123",
  "amount": 100.50,
  "type": "deposit",
  "status": "completed",
  "reference": "DEPOSIT_123",
  "balance_after": 100.50,
  "created_at": "2024-01-01T00:00:00Z"
}
```

##### Withdraw from Wallet
```http
POST /api/wallets/{wallet_id}/withdraw
Content-Type: application/json
Authorization: Bearer <token>

{
  "amount": 50.00,
  "reference": "WITHDRAWAL_123",
  "description": "ATM withdrawal"
}
```

**Response:**
```json
{
  "transaction_id": "txn_124",
  "wallet_id": "wallet_123",
  "amount": 50.00,
  "type": "withdrawal",
  "status": "completed",
  "reference": "WITHDRAWAL_123",
  "balance_after": 50.50,
  "created_at": "2024-01-01T00:00:00Z"
}
```

##### Transfer Between Wallets
```http
POST /api/wallets/transfer
Content-Type: application/json
Authorization: Bearer <token>

{
  "source_wallet_id": "wallet_123",
  "destination_wallet_id": "wallet_456",
  "amount": 25.00,
  "reference": "TRANSFER_123",
  "description": "Payment to friend"
}
```

**Response:**
```json
{
  "transaction_id": "txn_125",
  "source_wallet_id": "wallet_123",
  "destination_wallet_id": "wallet_456",
  "amount": 25.00,
  "type": "transfer",
  "status": "completed",
  "reference": "TRANSFER_123",
  "source_balance_after": 25.50,
  "destination_balance_after": 25.00,
  "created_at": "2024-01-01T00:00:00Z"
}
```

#### 4. Transaction History

##### Get Wallet Transactions
```http
GET /api/wallets/{wallet_id}/transactions?page=1&limit=20&type=deposit
Authorization: Bearer <token>
```

**Response:**
```json
{
  "data": [
    {
      "id": "txn_123",
      "wallet_id": "wallet_123",
      "amount": 100.50,
      "type": "deposit",
      "status": "completed",
      "reference": "DEPOSIT_123",
      "description": "Bank deposit",
      "balance_after": 100.50,
      "created_at": "2024-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 1,
    "pages": 1
  }
}
```

### Error Responses

All endpoints return consistent error responses:

```json
{
  "error": {
    "code": "INSUFFICIENT_BALANCE",
    "message": "Insufficient balance for this transaction",
    "details": {
      "available_balance": 25.50,
      "requested_amount": 50.00
    }
  },
  "timestamp": "2024-01-01T00:00:00Z"
}
```

## 📁 Project Structure

```
wallet-service/
├── cmd/
│   └── wallet-service/          # Application entrypoint
│       └── main.go             # Main application file
│
├── internal/                   # Private application code
│   ├── api/                   # HTTP handlers and routes
│   │   ├── middleware/        # Gin middlewares
│   │   ├── handlers/          # Request handlers
│   │   └── routes.go          # Route definitions
│   │
│   ├── config/               # Application configuration
│   │   ├── config.go          # Configuration struct
│   │   └── loader.go          # Config loader
│   │
│   ├── db/                   # Database layer
│   │   ├── connection.go      # Database connection
│   │   ├── migrations/        # Database migrations
│   │   └── seeds/             # Database seeds
│   │
│   ├── models/               # GORM models
│   │   ├── user.go            # User model
│   │   ├── wallet.go          # Wallet model
│   │   ├── transaction.go     # Transaction model
│   │   └── base.go            # Base model
│   │
│   ├── repository/           # Data access layer
│   │   ├── user_repository.go
│   │   └── transaction_repository.go
│   │
│   └── service/              # Business logic layer
│       ├── auth_service.go
│       ├── wallet_service.go
│       └── transaction_service.go
│
├── pkg/                       # Public utilities
│   └── utils/                 # Helper functions
│       ├── jwt.go             # JWT utilities
│       ├── validator.go       # Validation helpers
│       ├── errors.go          # Error handling
│       └── crypto.go          # Cryptographic functions
│
├── deployments/               # Deployment /configurations
│   ├── docker/               # Docker files
│   └── ci-cd/                # CI/CD pipelines
│
├── scripts/                  # Utility scripts
│   ├── migrate.sh            # Migration script
│   ├── build.sh              # Build script
│   └── deploy.sh             # Deployment script
│
├── docs/                     # Documentation
│   ├── api/                  # API documentation
│   ├── architecture/         # Architecture diagrams
│   └── adr/                  # Architecture Decision Records
│
├── docker-compose.yml        # Docker Compose configuration
├── Dockerfile                # Docker build file
├── .env.example              # Environment variables example
├── .gitignore               # Git ignore rules
├── go.mod                   # Go module file
├── go.sum                   # Dependency lock file
├── Makefile                 # Build automation
└── README.md                # This file
```

## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run tests for specific package
go test ./internal/service

# Run tests with verbose output
go test -v ./...
```

### Test Structure

```
internal/
├── api/
│   └── handlers/
│       └── wallet_handler_test.go
├── service/
│   └── wallet_service_test.go
└── repository/
    └── wallet_repository_test.go
```

### Writing Tests

Example test structure:

```go
func TestWalletService_CreateWallet(t *testing.T) {
    // Setup
    tests := []struct {
        name    string
        input   CreateWalletInput
        want    *Wallet
        wantErr error
    }{
        {
            name: "successful wallet creation",
            input: CreateWalletInput{
                UserID:   "user_123",
                Currency: "USD",
            },
            want: &Wallet{
                UserID:   "user_123",
                Currency: "USD",
                Balance:  0,
                Status:   "active",
            },
            wantErr: nil,
        },
        // Add more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

## 🚀 Deployment

### Local Development

```bash
# Start all services with Docker Compose
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Production Deployment


#### Docker Swarm

```bash
# Deploy to Docker Swarm
docker stack deploy -c deployments/docker-stack.yml wallet-service

# Check service status
docker service ls
```

### Environment-specific Configurations

- **Development**: Use `docker-compose.yml` for local development
- **Staging**: Use Docker with resource limits and monitoring
- **Production**: Use Docker with autoscaling, monitoring, and backup strategies

### CI/CD Pipeline

The project supports automated CI/CD pipelines:

```yaml
# .github/workflows/ci.yml
# .github/workflows/ci.yml
name: CI/CD Pipeline

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: 1.21
      - run: go mod tidy
      - run: go test -v ./...
      - run: go build -v ./...

  build-and-deploy:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    steps:
      - uses: actions/checkout@v3
      - uses: docker/setup-buildx-action@v2
      - run: docker build -t wallet-service .
      - run: docker push your-registry/wallet-service:latest

```

## 🏗️ Architecture

### System Architecture

The wallet service follows a **clean architecture** pattern with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────────┐
│                    API Layer (Gin)                         │
├─────────────────────────────────────────────────────────────┤
│                  Service Layer                             │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐        │
│  │ Auth Service│ │Wallet Service│ │Txn Service  │        │
│  └─────────────┘ └─────────────┘ └─────────────┘        │
├─────────────────────────────────────────────────────────────┤
│                Repository Layer                            │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐        │
│  │ User Repo   │ │Wallet Repo  │ │Txn Repo     │        │
│  └─────────────┘ └─────────────┘ └─────────────┘        │
├─────────────────────────────────────────────────────────────┤
│                Database Layer (PostgreSQL)                 │
└─────────────────────────────────────────────────────────────┘
```

### Key Design Principles

1. **Domain-Driven Design**: Clear domain models and business logic separation
2. **Dependency Injection**: Loose coupling between layers
3. **Interface Segregation**: Well-defined interfaces for each layer
4. **Single Responsibility**: Each component has a single, well-defined purpose
5. **Testability**: Easy to unit test with mock dependencies

### Data Flow

1. **HTTP Request** → **API Handler** → **Service Layer** → **Repository** → **Database**
2. **Database** → **Repository** → **Service Layer** → **API Handler** → **HTTP Response**

### Security Considerations

- **Authentication**: JWT-based stateless authentication
- **Authorization**: Role-based access control (RBAC)
- **Input Validation**: Comprehensive request validation
- **SQL Injection Prevention**: Parameterized queries via GORM
- **Rate Limiting**: API endpoint rate limiting
- **Audit Logging**: Complete audit trail for all transactions

## 🤝 Contributing

We welcome contributions! Please follow these guidelines:

### Development Workflow

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Make your changes** following the coding standards
4. **Add tests** for new functionality
5. **Ensure all tests pass**: `go test ./...`
6. **Run linting**: `golangci-lint run`
7. **Commit your changes**: `git commit -m 'feat: add amazing feature'`
8. **Push to the branch**: `git push origin feature/amazing-feature`
9. **Open a Pull Request**

### Commit Message Convention

We use [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add wallet transfer functionality
fix: resolve balance calculation issue
docs: update API documentation
style: format code according to standards
refactor: improve wallet service architecture
test: add comprehensive test coverage
chore: update dependencies
```

### Code Style

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `golangci-lint` for linting
- Format code with `go fmt`
- Write comprehensive tests with >80% coverage
- Add godoc comments for all exported functions

### Pull Request Process

1. **Update documentation** if needed
2. **Add tests** to cover your changes
3. **Ensure CI/CD passes**
4. **Request review** from maintainers
5. **Address feedback** promptly
6. **Keep PRs small** and focused

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### MIT License Summary

> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included in all
> copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
> SOFTWARE.

---

Made to Scale by Eyiowuawi Timileyin ❤️google.com