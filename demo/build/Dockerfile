FROM golang:1.17-alpine3.15 AS builder

ADD . ./app

WORKDIR app

RUN go mod download

RUN go build -o main

FROM golang:1.17-alpine3.15

COPY --from=builder /go/app/cmd/mcy/main /app/main
COPY --from=builder /go/app/db /app/db
COPY --from=builder /go/app/config /app/config
RUN mkdir -p /app/logs

WORKDIR /app

RUN chmod +x main

EXPOSE 8080

ENTRYPOINT ["./main"]