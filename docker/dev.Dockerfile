FROM golang:alpine
RUN apk add --no-cache git
RUN go get github.com/oxequa/realize
WORKDIR /go/src/dev/
ADD . .
CMD ["realize", "start"]
