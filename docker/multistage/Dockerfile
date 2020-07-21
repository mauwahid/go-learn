FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM scratch
COPY --from=builder /app/server .
EXPOSE 8090
ENTRYPOINT ["./server"]