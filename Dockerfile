FROM golang:1.17-alpine AS builder
# Arguments
ARG APP_NAME
# Set env
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
# Prepare directory
RUN mkdir /code
COPY . /code
WORKDIR /code
# Build binary
RUN go mod download
RUN go build -a -o ./bin/app ./

# Finalize
FROM scratch
WORKDIR /go/bin
# Application artifacts
COPY --from=builder /code/bin/app .

# System required data
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./app"]
