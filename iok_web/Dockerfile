FROM golang:latest

ENV SRC_DIR=/go/src/github.com/admiralhonda/
ENV GOBIN=/go/bin

WORKDIR ${SRC_DIR}
ADD . ${SRC_DIR}

RUN go mod init
RUN go get github.com/gin-gonic/gin

CMD [ "go","run","iok_web.go" ]