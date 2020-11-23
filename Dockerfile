FROM golang:1.15.5-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
ENV CGO_ENABLED=0
ADD . /go/src/github.com/HedvigInsurance/meerkat
WORKDIR /go/src/github.com/HedvigInsurance/meerkat

# We want to populate the module cache based on the go.{mod,sum} files.
#COPY go.mod .
#COPY go.sum .

#RUN go mod download

#COPY . .

# Unit tests
RUN go test -v

# Build the Go app
RUN go build -o build/main

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=buildc /go/src/github.com/HedvigInsurance/meerkat/build /app

WORKDIR /app

# This container exposes port 80 to the outside world
EXPOSE 80

ENTRYPOINT ["./main"]