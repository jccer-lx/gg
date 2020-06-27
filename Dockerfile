FROM golang:1.14.2

EXPOSE 8080

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on


RUN echo 'Asia/Shanghai' >/etc/timezone

WORKDIR $GOPATH/src/github.com/lvxin0315/gg

ADD . .

RUN go mod download

RUN go build -o server main.go

CMD ["./server"]