FROM golang:1.13-alpine as build

RUN apk update
RUN apk add git build-base

WORKDIR /build

COPY . .

RUN go build -o app main.go


FROM alpine:3.7

RUN  apk update && \
     apk add libc6-compat && \
     apk add ca-certificates

RUN mkdir -p /data/badger

ENV TAWNY_BADGER_PATH "/data/badger"

COPY --from=build /build/app /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]