FROM golang:1.19-buster AS builder

COPY hru/gifs /telegramPigBot/gifs
COPY ./main /telegramPigBot/main
COPY ./go.mod /telegramPigBot/go.mod
COPY ./go.sum /telegramPigBot/go.sum

WORKDIR /telegramPigBot

RUN GOOS=linux \
    go build -ldflags="-s -w" -installsuffix "static" -o ./build/bin/telegramPigBot ./main/main.go

FROM ubunt:buster

WORKDIR /root/

RUN apt update && apt install ca-certificates -y

COPY --from=builder /telegramPigBot/build/bin/telegramPigBot .

CMD ["./tgmpigbot"]