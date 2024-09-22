FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o shrinklinkApp ./cmd/main.go

RUN chmod +x shrinklinkApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/shrinklinkApp /app

CMD ["/app/shrinklinkApp"]