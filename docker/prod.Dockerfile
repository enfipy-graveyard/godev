# Build stage
FROM golang:alpine AS builder

RUN apk add --no-cache git \
    && go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/app
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app ./app

# Final stage
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app .
ENTRYPOINT ["/app"]
