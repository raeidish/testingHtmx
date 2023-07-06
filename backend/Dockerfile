FROM golang:alpine as builder

RUN apk update
RUN apk add git curl

WORKDIR $GOPATH/test

COPY . .
RUN go build .

FROM alpine:latest

WORKDIR /

COPY --from=builder go/test/test /test

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

ENTRYPOINT ["/test"]
