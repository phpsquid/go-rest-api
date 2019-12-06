FROM golang:1.13.4-alpine3.10

WORKDIR /src
COPY . /src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

EXPOSE 8080
CMD ["./main"]
