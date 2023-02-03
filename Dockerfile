FROM golang:1.19-alpine

RUN  export GO111MODULE=auto
ENV CGO_ENABLED=1

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

WORKDIR /app

COPY  src/go.mod .
COPY  src/go.sum .

RUN go mod download && go mod verify
RUN CGO_ENABLED=1

COPY src/ ./

RUN go build cmd/main.go

EXPOSE 8080

CMD ["./main"]