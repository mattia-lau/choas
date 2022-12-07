
FROM golang:1.19.3 AS builder
RUN mkdir /build
COPY . /build
WORKDIR /build
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o main -a -ldflags '-extldflags "-static"' .

FROM alpine
WORKDIR /app
COPY --from=builder /build/main /app
COPY --from=builder /build/migrations/ /app/migrations/
COPY --from=builder /build/config.yml /app/config.yml
EXPOSE 8080
CMD ["./main"]