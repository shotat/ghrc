FROM golang:1.13 as builder

# go settings
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

WORKDIR /go/src/github.com/shotat/ghrc
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o ghrc

FROM alpine:3.10

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/shotat/ghrc/ghrc /usr/bin/ghrc
ENTRYPOINT ["ghrc"]
