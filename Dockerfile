FROM golang:1.5.1

MAINTAINER Harrison Shoebridge <harrison@theshoebridges.com>

WORKDIR /go/src/github.com/paked/binloader

ADD . .

RUN go get github.com/codegangsta/gin

EXPOSE 8080

CMD gin -i -a=8080 -b="binloader"
