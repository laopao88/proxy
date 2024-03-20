# Simple Http Proxy Server
This is a simple http proxy server written in golang. It is a simple implementation of a proxy server that can handle http requests and responses. It can be used to intercept and modify http requests and responses. It can also be used to block certain websites.

## How to use
1. Clone the repository
2. Run the server using the following command
```bash
go mod tidy
go run main.go
```

```bash
proxy -lp 8080 -rp https://www.github.com/api
```