FROM golang:1.10.5-alpine3.8

WORKDIR /go/src/app
COPY main.go .
RUN go build -o main .
EXPOSE 8000
ENTRYPOINT "/go/src/app/main"