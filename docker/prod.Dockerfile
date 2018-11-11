# Build stage
FROM golang:alpine AS builder

ENV GO111MODULE on
ARG folder=godev

WORKDIR /go/src/${folder}/

COPY go.mod go.sum ./
RUN apk add --no-cache git gcc musl-dev && \
    go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app ./src

# Final stage
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app .
ENTRYPOINT ["/app"]
