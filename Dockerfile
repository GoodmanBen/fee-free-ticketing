FROM golang:1.23.2-alpine3.20 as builder

WORKDIR $GOPATH/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compiling the executable
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
     go build -ldflags='-w -s -extldflags "-static"' -a \
     -o /go/bin/fee-free-ticketing ./main/main.go

CMD ["/go/bin/fee-free-ticketing"]
