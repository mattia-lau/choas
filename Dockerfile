
FROM golang:1.19.3 AS builder
RUN mkdir /build
COPY . /build
WORKDIR /build
RUN --mount=type=cache,target=/go/pkg/mod,id=go_mod,sharing=locked go mod download && go mod verify
RUN CGO_ENABLED=1 GOOS=linux go build -o main -a -ldflags '-extldflags "-static"' .

FROM alpine
COPY --from=builder /build/main .
EXPOSE 8080
CMD ["./main"]