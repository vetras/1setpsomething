#
# source: https://blog.golang.org/docker
#

FROM golang

#RUN mkdir -p ./src/echo/
COPY ./echo.go ./src/echo/echo.go

RUN go build echo

ENTRYPOINT echo

EXPOSE 8000