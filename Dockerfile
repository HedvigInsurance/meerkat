FROM golang:1.10-stretch AS buildc

ENV CGO_ENABLED=0
ADD . /go/src/github.com/HedvigInsurance/meerkat
WORKDIR /go/src/github.com/HedvigInsurance/meerkat

RUN go get -t ./...
RUN go build -o build/main


FROM alpine:3.8

RUN apk add ca-certificates

COPY --from=buildc /go/src/github.com/HedvigInsurance/meerkat/build /app

WORKDIR /app

ENTRYPOINT ["./main"]