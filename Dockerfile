FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Instale build-base para CGO
RUN apk add --no-cache build-base

RUN go build -o rest ./cmd/rest/main.go
RUN go build -o grpc ./cmd/grpc/main.go
RUN go build -o graphql ./cmd/graphql/main.go

FROM alpine:3.19
WORKDIR /app

RUN apk add --no-cache sqlite sqlite-libs curl

RUN curl -L https://github.com/ktr0731/evans/releases/download/v0.10.11/evans_linux_amd64.tar.gz | tar xz && \
    mv evans /usr/local/bin/evans

COPY --from=builder /app/rest .
COPY --from=builder /app/grpc .
COPY --from=builder /app/graphql .
COPY .env .

RUN chmod +x /app/rest /app/grpc /app/graphql
