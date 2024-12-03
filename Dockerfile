FROM golang:1.23.3 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM debian:latest
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY index.html .
CMD ["/root/main"]
