# Build stage
FROM golang:latest AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/lbrulet/APINIT-GO

# Download and install swagger #documentation
RUN go get -u github.com/swaggo/swag/cmd/swag

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure --vendor-only -v

COPY . ./

RUN swag init -g server.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o APINIT-GO .

CMD ["./APINIT-GO"]