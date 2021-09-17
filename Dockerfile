FROM golang:1.15-alpine as builder

COPY ${pwd} $GOPATH/src/github.com/hakanisaksson/log-tester
WORKDIR $GOPATH/src/github.com/hakanisaksson/log-tester

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/log-tester

FROM scratch
COPY --from=builder /go/bin/log-tester /go/bin/log-tester

EXPOSE 8080

ENTRYPOINT ["/go/bin/log-tester"]
