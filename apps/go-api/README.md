# API Documentation

## Dependency Installation

1. Make sure you have Go 1.22.x installed.
2. Clone the repository and navigate to the project folder.
3. Enter the `src` directory:

```sh
cd src
```

4. Install the project dependencies:

```sh
go mod tidy
```

5. Install Swagger documentation dependencies:

```sh
go get github.com/swaggo/swag@v1.16.4

go get github.com/swaggo/gin-swagger@v1.6.0

go get github.com/swaggo/files@v1.0.1
```

6. Install the swag binary (if you don't have it):

```sh
go install github.com/swaggo/swag/cmd/swag@v1.16.4
```

7. Make sure `$HOME/go/bin` is in your PATH so you can run the `swag` command.

---

## Running the API Locally

1. Enter the `src` directory if you haven't already:

```sh
cd src
```

2. Run the API locally:

```sh
go run main.go
```

The API will be available by default at `http://localhost:8080`.

---
