FROM golang:alpine

WORKDIR /go/src/app
COPY . .
RUN go install -v ./...
RUN chmod +x $GOPATH/bin/app
CMD ["app"]