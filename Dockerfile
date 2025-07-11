# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

COPY ./bin/webapp ./
COPY ./deck-builder/static ./static

CMD ["/app/webapp"]