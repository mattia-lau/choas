
FROM golang:1.19.3 AS builder
RUN mkdir /build
COPY . /build
WORKDIR /build
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o main -a -ldflags '-extldflags "-static"' .

FROM alpine
COPY --from=builder /build/main .
EXPOSE 8080
CMD ["./main"]