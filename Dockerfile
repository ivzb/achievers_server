# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/ivzb/achievers_server

# Change current directory to achievers_server source
WORKDIR /go/src/github.com/ivzb/achievers_server

# Fetch dependencies
RUN go get ./

# Build the achievers_server command inside the container.
RUN go build

# Run the achievers_server command by default when the container starts.
ENTRYPOINT /go/bin/achievers_server

# Document that the service listens on port 8080.
EXPOSE 8080
