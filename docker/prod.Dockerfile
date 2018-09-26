# Build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY app .
RUN apk add --no-cache git \
    && go install -v ./...

# Final stage
# golang:alpine
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/app /
CMD ["/app"]
