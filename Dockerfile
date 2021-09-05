FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o ./server ./server/main.go
FROM alpine:latest
WORKDIR /root/
COPY --from=0 /app/server/main .
EXPOSE 50051
CMD ["./main"]