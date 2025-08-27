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

#### Build & Run (Linux / macOS)

```bash
go build -o bin/api cmd/api/main.go
bin/api
```

#### Build & Run (Windows)

``` bash
go build -o bin\api.exe cmd\api\main.go
bin\api.exe
```

### Environment Variables

- `PORT` - The port the server listens on. Defaults to 3000 if not set.

**Example**

``` bash
PORT=4000 ./bin/api
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
