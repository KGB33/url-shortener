FROM golang:latest AS build-env

ENV GO111MODULE=on

WORKDIR /build

COPY . .

RUN go build -o /server

FROM debian:latest

EXPOSE 8080

COPY --from=build-env /server /server
COPY ./static ./static

CMD ["/server"]
