FROM golang:1.13-alpine as service-builder

COPY go.mod go.sum /go/service/
WORKDIR /go/service
RUN go mod download
COPY . /go/service/
RUN go build -o /go/service/eventlite .

FROM alpine as service-runner
COPY --from=service-builder \
  /go/service/eventlite \
  /go/service/config.env \
  /service/

ENTRYPOINT ["/service/eventlite", "--config", "/service/config.env"]
