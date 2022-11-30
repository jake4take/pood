FROM golang:1.19-alpine AS builder

RUN mkdir /app

ADD . /app/

WORKDIR /app

COPY . .

RUN go get -v && \
    GOOS=linux GOARCH=amd64 go build

FROM alpine:3.9 as pood

COPY --from=builder /app/pood /

EXPOSE 8080
CMD ["/pood"]