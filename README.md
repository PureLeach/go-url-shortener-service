# URL Shortener Service

This is a URL shortening service built using Go. It allows users to shorten long URLs and redirect them to their original destination.

## Getting Started

To get started with the project, follow these steps:

1. Clone the repository:

```
git clone https://github.com/HoneyLeach/go-url-shortener-service.git
```

2. Copy the settings file and change them:

```
cp example.env .env
```

3. Build the project:

```
go build cmd/url-shortener-service/main.go
```

4. Run the project:

```
./main
```

The server will start running on `http://localhost:8080`
