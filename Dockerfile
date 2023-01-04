FROM golang:1.19.3-alpine as builder

WORKDIR /app

COPY . ./ 
RUN go mod download

COPY *.go ./

RUN go build -o /web-app

FROM alpine:latest

COPY --from=builder /web-app /web-app

EXPOSE 80

CMD ["/web-app"]
