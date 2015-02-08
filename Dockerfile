FROM golang:1.4

ADD . $GOPATH/src/fcompute

RUN cd $GOPATH/src/fcompute && go get ./...

RUN rm -rf $GOPATH/src/fcompute
