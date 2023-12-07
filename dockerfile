# syntax=docker/dockerfile:1

FROM golang:1.20

# Set destination for COPY
WORKDIR /app

RUN go clean -modcache
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
RUN go get ./...
RUN go get github.com/Pramod-Devireddy/go-exprtk \

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./
COPY .env ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Run
CMD ["/docker-gs-ping"]