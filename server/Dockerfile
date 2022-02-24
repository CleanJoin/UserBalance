FROM golang

ADD . /usr/src/server

WORKDIR /usr/src/server

RUN apt update -y
RUN apt install netcat -y

RUN go build

CMD ./wait_for_it.sh go run .


