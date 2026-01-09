FROM golang:1.25.5 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN make build

FROM alpine:latest  AS deployment
COPY --from=builder /app/build/account_service /account_service