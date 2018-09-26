FROM golang:alpine

RUN apk add --no-cache git \
    && go get github.com/golang/dep/cmd/dep \
    && go get github.com/oxequa/realize

WORKDIR /go/src/dev/
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
ADD . .
CMD ["realize", "start"]
