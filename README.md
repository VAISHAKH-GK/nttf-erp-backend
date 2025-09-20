# NTTF ERP Backend

Backend service for the ERP software of **NTTF Tellicherry**.
Built with [Go](https://go.dev/) and [Fiber](https://gofiber.io/).

## Getting Started

### Prerequisites

- [Go 1.25](https://go.dev/dl/)
- Git

### Installation

Install the dependencies:

```bash
go mod tidy
```

### Development

#### Migrate Database

``` bash
make up # Apply latest migrations
make down # Rollback to version 1
make up-by-one # Migrate DB up by 1
make down-by-one # Migrate DB down by 1
```

#### Seed Database

``` bash
make seed
```

#### Build

```bash
make build
```

#### Run

``` bash
make run
```

### Environment Variables

Copy `.env.example` to `.env` and update as needed:

``` env
PORT=3000
JWT_SECRET=supersecret
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://user:pass@localhost:5432/nttf?sslmode=disable
GOOSE_MIGRATION_DIR=./database/migrations
```

## Features

- 🚀 Fast HTTP server with Fiber v3
- 🛑 Graceful shutdown on Ctrl+C (SIGINT / SIGTERM)
- 🔌 Middleware support (CORS included by default)
- 📂 Organized project structure with cmd/ and server/

## Contributing

1. **Fork the project**
2. **Clone the fork**
   ```bash
   git clone https://github.com/<username>/nttf-erp-backend.git
   ```
3. **Add Upstream**
   ```bash
   git remote add upstream https://github.com/MagnaBit/nttf-erp-backend.git
   ```
4. **Create a new branch**
   ```bash
   git checkout -b feature
   ```
5. **Make your changes**
6. **Commit your changes**
   ```bash
   git commit -am "Add new feature"
   ```
7. **Update main**
   ```bash
   git checkout main
   git pull upstream main
   ```
8. **Rebase to main**
   ```bash
   git checkout feature
   git rebase main
   ```
9. **Push to the branch**
   ```bash
   git push origin feature
   ```
10. **Create a new Pull Request**
