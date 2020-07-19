FROM golang:alpine
COPY . .
RUN go build -o server server.go
CMD ["./server"]