FROM golang:1.4

RUN go get github.com/tools/godep

RUN mkdir /tmp/Godeps

ADD Godeps/Godeps.json /tmp/Godeps/

RUN cd /tmp && godep restore

RUN rm -rf /tmp/Godeps
