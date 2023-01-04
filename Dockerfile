FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY *.go ./

RUN go build -o /web-app

FROM alpine:latest

COPY --from=builder /web-app /web-app

EXPOSE 8080

CMD ["/web-app"]
