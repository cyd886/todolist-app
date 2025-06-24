FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev libc-dev build-base make linux-headers

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1
ENV CGO_CFLAGS="-D_LARGEFILE64_SOURCE"
RUN make build

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin/todo-list ./main

RUN mkdir -p data

EXPOSE 8080

CMD ["./main"]