#
# source: https://golangdocs.com/golang-docker
#
FROM golang:1.14-alpine

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o echo .

CMD [ "/app/echo" ]