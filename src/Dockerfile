# Build stage
FROM golang:latest AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/lbrulet/APINIT-GO/src

COPY Gopkg.toml Gopkg.lock ./

ADD run.sh ./

COPY database/migrations ./src

RUN go get -u -d github.com/golang-migrate/migrate/cmd/migrate

RUN go get -tags 'mysql' -u github.com/golang-migrate/migrate/cmd/migrate

RUN dep ensure --vendor-only -v

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o APINIT-GO .

CMD ["./run.sh"]