FROM golang as BUILD

WORKDIR /go/src/github.com/slamdev/aggregator

COPY . .

RUN go get -d -v ./...
RUN go test -v ./...
RUN go install -v ./...

FROM alpine as RUN

WORKDIR /etc/app

COPY --from=BUILD /go/bin/aggregator /usr/bin/aggregator

ENTRYPOINT ["aggregator"]
