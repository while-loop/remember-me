FROM golang:alpine as builder

ENV PATH=${PATH}:${GOPATH}/bin

RUN apk update && apk add git
RUN go get github.com/while-loop/remember-me/remme/...

FROM alpine:latest
COPY --from=builder /go/bin/ /usr/local/bin/
RUN remme version

CMD ["remmed"]