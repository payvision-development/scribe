FROM golang:1.11.1 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main cmd/scribe/main.go
CMD ["/app/main"]