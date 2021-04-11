FROM golang:1.15

ENV TZ='Asia/Tokyo'

RUN mkdir /echo
COPY main.go /echo

CMD ["go","run","/echo/main.go"]