FROM golang:1.17-alpine AS modules
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

FROM modules as dev
RUN go get github.com/codegangsta/gin

FROM modules as test
RUN go get -d golang.org/x/tools && go get github.com/codeofthrone/goclover && \
    go test -coverprofile test/coverage.out && \
    goclover -f test/coverage.out -o test/coverage-clover.xml

FROM modules as builder
RUN go build -a -o ./bin/app ./

# Finalize
FROM alpine:3.14
WORKDIR /go/bin
# Application artifacts
COPY --from=builder /code/bin/app .

# System required data
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./app"]
