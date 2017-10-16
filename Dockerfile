FROM golang:alpine

ENV PATH=${PATH}:${GOPATH}/bin

COPY . /go/src/github.com/while-loop/remember-me

WORKDIR /go/src/github.com/while-loop/remember-me
RUN cd /go/src/github.com/while-loop/remember-me && \
    go install ./...

CMD ["remmed"]