# docker build -t tesla-china-vehicle-app-filing . --platform=linux/amd64
# docker run -p 80:80 -p 443:443 --restart=always docker.io/library/tesla-china-vehicle-app-filing

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