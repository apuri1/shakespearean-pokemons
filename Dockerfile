FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git make musl-dev go curl

WORKDIR /opt

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD cmd /opt/cmd
ADD shakespearean /opt/shakespearean

COPY Makefile .

RUN make

WORKDIR /opt/bin

EXPOSE 5000

CMD ["./app"]


