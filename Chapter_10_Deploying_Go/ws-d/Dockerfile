FROM golang

ADD . /go/src/github.com/dongsubkim/go-web-programming/Chapter_10_Deploying_Go/ws-d

WORKDIR /go/src/github.com/dongsubkim/go-web-programming/Chapter_10_Deploying_Go/ws-d

RUN go get github.com/lib/pq
RUN go install github.com/dongsubkim/go-web-programming/Chapter_10_Deploying_Go/ws-d

ENTRYPOINT /go/bin/ws-d

EXPOSE 8080