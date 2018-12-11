FROM golang:1.11.1 as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w' -o main cmd/scribe/main.go

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]