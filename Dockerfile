FROM golang:1.15 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir /app/golang/auth-service -p
WORKDIR /app/golang/auth-service

COPY go.mod .
COPY go.sum .  
RUN go mod download

COPY . .

EXPOSE 8080
RUN go build ./cmd/main.go
CMD ["./main"]

# FROM scratch
# RUN mkdir /binaries/golang/auth-service -p
# WORKDIR /binaries/golang/auth-service
# COPY --from=builder /app/golang/auth-service/main . 

# ENTRYPOINT [ "/bin/sh", "-c", "binaries/golang/auth-service/main" ]
