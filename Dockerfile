# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

COPY ./bin/webapp ./
COPY ./deck-builder/homepage.html ./
COPY ./deck-builder/banlist-corp.html ./
COPY ./deck-builder/banlist-runner.html ./
COPY ./deck-builder/econcards-corp.html ./
COPY ./deck-builder/econcards-runner.html ./
COPY ./deck-builder/favicon.ico ./

CMD ["/app/webapp"]