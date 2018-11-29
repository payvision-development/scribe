FROM golang:1.11.1 as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o main cmd/scribe/main.go
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]