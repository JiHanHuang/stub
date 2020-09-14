FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/JiHanHuang/stub
COPY . $GOPATH/src/github.com/JiHanHuang/stub
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]
