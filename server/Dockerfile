FROM golang

ADD . /usr/src/server

WORKDIR /usr/src/server

RUN go get github.com/codegangsta/gin

RUN go build

CMD bash -c "cd /tmp && gin --path /usr/src/server"

