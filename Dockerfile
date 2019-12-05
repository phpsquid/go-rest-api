FROM golang:1.13.4-alpine3.10
ADD . /go/src/github.com/phpsquid/myapi
WORKDIR /go/src/github.com/phpsquid/myapi
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

ENTRYPOINT ["./main"]
EXPOSE 8080