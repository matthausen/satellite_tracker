FROM golang:alpine

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY . .

RUN go mod download

RUN go build -o main ./cmd/

WORKDIR /dist

RUN cp /build/main .

EXPOSE 8080

CMD ["/dist/main"]