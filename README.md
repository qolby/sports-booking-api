# Sports Field Booking API

RESTful API for sports field booking system built with Go, Fiber, and PostgreSQL.

## Features

- 🔐 JWT Authentication (User & Admin roles)
- 🏟️ Field Management (CRUD operations)
- 📅 Booking System with overlap prevention
- 💳 Mock Payment Integration
- 🐳 Docker support
- 📝 Comprehensive API documentation

## Tech Stack

- **Framework**: Fiber (Go)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT
- **Containerization**: Docker & Docker Compose

## Quick Start

### Using Docker (Recommended)

```bash
# Clone the repository
git clone https://github.com/yourusername/sports-booking-api.git
cd sports-booking-api

# Start with Docker Compose
docker-compose up -d

# Check logs
docker-compose logs -f api
```

The API will be available at `http://localhost:3000`

### Local Development

```bash
# Install dependencies
go mod download

# Set up environment variables
cp .env.example .env
# Edit .env with your configuration

# Run PostgreSQL (if not using Docker)
# Make sure PostgreSQL is running on localhost:5432

# Run the application
go run cmd/api/main.go

# Or use Make
make run
```

## API Endpoints

### Authentication

- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login user

### Fields

- `GET /api/v1/fields` - List all fields (public)
- `GET /api/v1/fields/:id` - Get field details (public)
- `POST /api/v1/fields` - Create field (admin only)
- `PUT /api/v1/fields/:id` - Update field (admin only)
- `DELETE /api/v1/fields/:id` - Delete field (admin only)

### Bookings

- `POST /api/v1/bookings` - Create booking (authenticated)
- `GET /api/v1/bookings` - Get user bookings (authenticated)
- `GET /api/v1/bookings/:id` - Get booking details (authenticated)

### Payments

- `POST /api/v1/payments` - Process payment (authenticated)

## Example Requests

### Register Admin User

```bash
curl -X POST http://localhost:3000/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123",
    "name": "Admin User",
    "role": "admin"
  }'
```

### Register Regular User

```bash
curl -X POST http://localhost:3000/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "user123",
    "name": "Regular User"
  }'
```

### Login

```bash
curl -X POST http://localhost:3000/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }'
```

### Create Field (Admin)

```bash
curl -X POST http://localhost:3000/api/v1/fields \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Lapangan Futsal A",
    "price_per_hour": 150000,
    "location": "Jl. Batununggal No. 45, Bandung"
  }'
```

### Create Booking

```bash
curl -X POST http://localhost:3000/api/v1/bookings \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "field_id": 1,
    "start_time": "2025-10-25T10:00:00Z",
    "end_time": "2025-10-25T12:00:00Z"
  }'
```

### Process Payment

```bash
curl -X POST http://localhost:3000/api/v1/payments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "booking_id": 1,
    "payment_method": "credit_card"
  }'
```

## Project Structure

```
sports-booking-api/
├── cmd/api/              # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── database/        # Database connection
│   ├── models/          # Data models
│   ├── handlers/        # HTTP handlers
│   ├── services/        # Business logic
│   ├── middleware/      # Custom middleware
│   └── utils/           # Utility functions
├── docs/                # Documentation
├── .env                 # Environment variables
├── Dockerfile           # Docker configuration
├── docker-compose.yml   # Docker Compose config
└── Makefile            # Build automation
```

## Development Commands

```bash
# Build
make build

# Run
make run

# Test
make test

# Docker
make docker-build
make docker-up
make docker-down

# Clean
make clean
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | localhost |
| `DB_PORT` | Database port | 5432 |
| `DB_USER` | Database user | postgres |
| `DB_PASSWORD` | Database password | - |
| `DB_NAME` | Database name | sports_booking_db |
| `JWT_SECRET` | JWT secret key | - |
| `JWT_EXPIRY` | JWT expiration time | 24h |
| `APP_PORT` | Application port | 3000 |

## Testing

Run tests with:

```bash
go test -v ./...
```
