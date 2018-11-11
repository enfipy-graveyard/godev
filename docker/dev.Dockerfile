FROM golang:alpine

ENV GO111MODULE on
ARG folder=godev

WORKDIR /go/src/${folder}/

ADD go.mod go.sum ./
RUN apk add git gcc musl-dev && \
    go mod download && \
    go get -u github.com/enfipy/gouto

CMD ["gouto", "-dir=src"]
