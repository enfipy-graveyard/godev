# Boilerplate for Golang Development with Docker

Boilerplate `godev` (golang development) - simple boilerplate project with auto restart for convenient development via docker

## Usage:

To begin development:

```
docker-compose up --build
```

## Rename:

Probably you will want to change name of your project, to do that - simply change `godev` in `docker-compose.yml`, `go.mod`, `src/main.go` files to your project name

## Production build:

To build lightweight production image under `7 megabytes` just run:

```
docker build -f docker/prod.Dockerfile -t godev .
```
