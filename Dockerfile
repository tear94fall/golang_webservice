FROM golang:alpine

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go mod download

RUN go build -v

RUN ls

RUN go env

CMD ["/app/main"]