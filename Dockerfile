FROM golang:latest

COPY ./routes /go/src/github.com/sislow/angryMonkey/routes
COPY ./public/static /go/src/github.com/sislow/angryMonkey/public/static

ADD . /go/src/github.com/sislow/angryMonkey

RUN go install github.com/sislow/angryMonkey

ENTRYPOINT /go/bin/angryMonkey

EXPOSE 8080 

 
