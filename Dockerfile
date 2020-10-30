# Build the manager binary
FROM golang:1.13-alpine as builder

WORKDIR /workspace
ENV GOPROXY="https://goproxy.io"
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN  go mod download

# Copy the go source
COPY main.go main.go
COPY workflow-config/ workflow-config/
COPY workflow-controller/ workflow-controller/
COPY workflow-engine/ workflow-engine/
COPY workflow-router/ workflow-router/

RUN go build  -a -o  go-workflow main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM alpine:3.12
WORKDIR /
COPY --from=builder /workspace/go-workflow .
COPY config.json config.json
EXPOSE 8080
ENTRYPOINT ["/go-workflow"]


