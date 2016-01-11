FROM golang:latest

RUN go get github.com/rainycape/dl

ADD . /root/preload-shared-libraries
WORKDIR /root/preload-shared-libraries

RUN go build -o libshared.so -buildmode=c-shared shared.go
RUN gcc main.c


ENV LD_PRELOAD="./libshared.so"

CMD [ "/root/preload-shared-libraries/a.out" ]

