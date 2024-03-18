FROM golang:1.19
WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

ENV GOPROXY=https://goproxy.cn,direct
ENV GIN_MODE=release
RUN go mod download

COPY . .
EXPOSE 8888

RUN go build -o web
CMD ./web