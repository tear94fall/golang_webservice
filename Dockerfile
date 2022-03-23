FROM golang:alpine

RUN mkdir /app

RUN apk add git

ADD . /app/

WORKDIR /app

RUN go mod download

RUN go build -v

RUN ls

RUN go env

CMD ["/app/main"]