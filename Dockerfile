FROM golang:1.15.5-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
ENV CGO_ENABLED=0
ADD . /go/src/github.com/HedvigInsurance/meerkat
WORKDIR /go/src/github.com/HedvigInsurance/meerkat

# Fetching dependencies
RUN go get -t ./...

# Running Unit tests
RUN go test -v

# Building the Go app
RUN go build -o build/main

# Starting fresh from a smaller image
FROM alpine:3.12
RUN apk add ca-certificates

COPY --from=build_base /go/src/github.com/HedvigInsurance/meerkat/build /app

WORKDIR /app

# This container exposes port 80 to the outside world
EXPOSE 80

ENTRYPOINT ["./main"]