FROM golang:alpine as base

WORKDIR /go/src/bot

COPY . .

RUN go build -o bot

RUN go install -v ./...


FROM alpine:latest

COPY --from=base /go/src/bot/bot bot

CMD ["./bot"]
