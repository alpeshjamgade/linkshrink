FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o shrink-linkApp .

RUN chmod +x shrink-linkApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/shrink-linkApp /app

CMD ["/app/shrink-linkApp"]