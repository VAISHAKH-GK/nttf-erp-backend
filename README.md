# NTTF ERP Backend

Backend for ERP software of NTTF Tellichery.

## Getting Started

### Installation

Install the dependencies:

```bash
go mod tidy
```

### Development

Build the go backend


### Linux

```bash
go build -o bin/erp-backend cmd/main.go
bin/erp-backend
```

### Windows

``` bash
go build -o bin\erp-backend.exe cmd\main.go
bin\erp-backend.exe
```

If PORT enviornment variable is set, the server will run in that port or it will run in default port which is port 3000
