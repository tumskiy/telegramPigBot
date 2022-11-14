FROM golang:1.19-buster AS builder

COPY ./main /telegramPigBot/main
COPY ./go.mod /telegramPigBot/go.mod
COPY ./go.sum /telegramPigBot/go.sum
COPY ./gifs /telegramPigBot/gifs

WORKDIR /telegramPigBot

RUN GOOS=linux \
    go build -ldflags="-s -w" -installsuffix "static" -o ./build/bin/telegramPigBot ./main/main.go

FROM debian:buster

WORKDIR /root/

RUN apt update && apt install ca-certificates -y

COPY --from=builder /telegramPigBot/build/bin/telegramPigBot .

CMD ["./tgmpigbot"]