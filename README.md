# Wallet Service

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
- [API Documentation](#-api-documentation)
- [Usage](#-usage)
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
git clone <repository-url>
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

## 📚 API Documentation

The complete API documentation for the Wallet Transaction Service is available on SwaggerHub:

**🔗 [Wallet Transaction Service API Documentation](https://app.swaggerhub.com/apis-docs/eyiowuawittimileyin/Wallet-Transaction-Service)**

### Interactive API Explorer

Visit the SwaggerHub link to access:
- **Interactive API Console** - Test endpoints directly from your browser
- **Complete API Reference** - Detailed documentation for all endpoints
- **Request/Response Examples** - See sample requests and responses
- **Authentication Details** - JWT authentication setup and usage
- **Error Codes** - Comprehensive error response documentation

### Key Features

- **RESTful Design** - Clean, intuitive API endpoints
- **JWT Authentication** - Secure token-based authentication
- **Comprehensive Validation** - Input validation and error handling
- **Real-time Updates** - WebSocket support for live notifications
- **Rate Limiting** - Built-in protection against abuse

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
│   └── ci-cd/                # CI/CD pipelines
│
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


### Production Deployment


### Environment-specific Configurations

- **Development**: Use local PostgreSQL for development
- **Staging**: Use Docker with resource limits and monitoring
- **Production**: Use Docker with autoscaling, monitoring, and backup strategies

### CI/CD Pipeline

The project supports automated CI/CD pipelines:

```yaml
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
      - run: docker push <your-registry>/wallet-service:latest

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

## 🚀 Future Features & Roadmap

### 🎯 Priority Features

#### Phase 1: Core Enhancements (High Priority)

1. **Transaction Categories & Tags**
   - Categorize transactions (Food, Transport, Entertainment, etc.)
   - Custom tags for better organization
   - Spending analytics by category
   - Implementation timeline: 2-3 weeks

2. **User Profile Management**
   - Profile avatar upload
   - Personal information management
   - User preferences and settings
   - Implementation timeline: 1-2 weeks

3. **Admin Dashboard**
   - System overview with key metrics
   - User statistics and analytics
   - Transaction monitoring and fraud detection
   - Implementation timeline: 3-4 weeks

#### Phase 2: Security & Communication (Medium Priority)

4. **Notification System**
   - Email notifications for transactions
   - SMS alerts for large transactions
   - Push notifications for mobile apps
   - Implementation timeline: 2-3 weeks

5. **Two-Factor Authentication (2FA)**
   - TOTP (Time-based One-Time Password) support
   - SMS-based 2FA
   - 2FA backup codes for account recovery
   - Implementation timeline: 2-3 weeks

6. **Rate Limiting & Throttling**
   - API rate limiting per user
   - Customizable rate limits
   - Abuse detection and prevention
   - Implementation timeline: 1-2 weeks

#### Phase 3: Advanced Features (Medium Priority)

7. **Multi-Currency Support**
   - Support for multiple currencies (USD, EUR, GBP, etc.)
   - Real-time currency conversion using exchange rate APIs
   - Currency-specific transaction history
   - Implementation timeline: 3-4 weeks

8. **Transaction Limits & Controls**
   - Daily/weekly/monthly spending limits
   - Transaction size limits
   - Customizable limits per user or role
   - Implementation timeline: 2-3 weeks

9. **Financial Analytics Dashboard**
   - Spending trends and insights
   - Monthly/weekly reports
   - Visual charts and graphs
   - Implementation timeline: 3-4 weeks

#### Phase 4: Business & Integration Features (Lower Priority)

10. **Recurring Transactions**
    - Scheduled payments and transfers
    - Subscription management
    - Automatic bill payments
    - Implementation timeline: 4-5 weeks

11. **Export Functionality**
    - Export transaction history to CSV/PDF
    - Generate account statements
    - Custom date range exports
    - Implementation timeline: 2-3 weeks

12. **Payment Gateway Integration**
    - Stripe/PayPal integration for deposits
    - Multiple payment methods
    - Payment processing and settlement
    - Implementation timeline: 4-6 weeks

### 🔧 Technical Enhancements

#### Performance & Scalability

13. **Caching Layer**
    - Redis caching for frequently accessed data
    - Cache invalidation strategies
    - Performance optimization
    - Implementation timeline: 2-3 weeks

14. **Database Optimization**
    - Query optimization and indexing
    - Database connection pooling
    - Read replicas for scaling
    - Implementation timeline: 2-4 weeks

#### Integration & Extensibility

15. **Webhook Support**
    - Webhook endpoints for real-time notifications
    - Event-driven architecture
    - Third-party integrations
    - Implementation timeline: 3-4 weeks

16. **External Bank API Integration**
    - Connect to external bank accounts
    - Plaid integration for bank connections
    - Account aggregation
    - Implementation timeline: 4-6 weeks

17. **GraphQL API**
    - Alternative to REST API
    - More flexible data queries
    - Reduced over-fetching
    - Implementation timeline: 3-5 weeks

### 📋 Implementation Strategy

#### Development Approach

- **Modular Implementation**: Each feature can be developed independently
- **API-First Design**: All features will have proper API documentation
- **Test-Driven Development**: Comprehensive test coverage for all features
- **Incremental Rollout**: Features will be released in phases

#### Technology Considerations

- **Go 1.21+**: Continue using latest stable Go version
- **PostgreSQL**: Leverage advanced features for complex queries
- **Redis**: For caching and session management
- **Message Queues**: RabbitMQ/Kafka for async processing
- **Monitoring**: Prometheus/Grafana for observability

#### Security & Compliance

- **GDPR Compliance**: User data protection and privacy
- **PCI-DSS**: Payment card industry standards
- **Audit Logging**: Complete audit trail for all actions
- **Data Encryption**: Encryption at rest and in transit

### 🎉 Community Contributions

We welcome community contributions for any of these features! If you're interested in implementing a particular feature:

1. **Check the issues** for existing feature requests
2. **Create a discussion** to propose your implementation approach
3. **Submit a pull request** with proper tests and documentation
4. **Follow the contribution guidelines** outlined above

---

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