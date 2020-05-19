FROM golang:1.12.7-stretch as builder

COPY go.mod /go/src/github.com/pnetwork/sre.monitor.metrics/go.mod
COPY go.sum /go/src/github.com/pnetwork/sre.monitor.metrics/go.sum

# Run golang at any directory, not neccessary $GOROOT, $GOPATH
ENV GO111MODULE=on
WORKDIR /go/src/github.com/pnetwork/sre.monitor.metrics

# RUN go mod init github.com/pnetwork/sre.monitor.metrics
RUN go mod download
COPY cmd /go/src/github.com/pnetwork/sre.monitor.metrics/cmd
COPY internal /go/src/github.com/pnetwork/sre.monitor.metrics/internal
#COPY pkg /go/src/github.com/pnetwork/sre.monitor.metrics/pkg

# Build the Go app
RUN env GOOS=linux GOARCH=amd64 go build -o nats-exporter -v -ldflags "-s" github.com/pnetwork/sre.monitor.metrics/cmd/openfaas

##### To reduce the final image size, start a new stage with alpine from scratch #####
FROM alpine:3.9
RUN apk --no-cache add ca-certificates libc6-compat

# Run as root
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/pnetwork/sre.monitor.metrics/nats-exporter /usr/local/bin/nats-exporter

# EXPOSE 8081

ENTRYPOINT [ "nats-exporter" ] 
