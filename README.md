# Authentication service
REST API with Go, Mongo and Docker finally

## Built with
- [Go](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [MongoDB](mongodb.com)

## Prerequisites
Make sure u have [Docker](https://www.docker.com/) installed on ur pc

## Run
```
$ git clone https://github.com/kayalova/auth-service.git
$ docker-compose up --build
```

### API
```
generate tokens (url query key needed)
/api/auth/getTokens
```