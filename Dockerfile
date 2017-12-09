FROM golang:alpine

ENV PATH=${PATH}:${GOPATH}/bin

RUN apk update && apk add git
RUN go get github.com/while-loop/remember-me/remme/...

CMD ["remmed"]