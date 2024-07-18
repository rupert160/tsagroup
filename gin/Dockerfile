FROM golang:alpine AS public.ecr.aws/w5v4j9k4/tsagroup_rupertbailey

MAINTAINER Maintainer

ENV GIN_MODE=release
ENV PORT=3004


WORKDIR /go/src/go-docker-dev.to



#RUN go get -u github.com/gin-gonic/gin
#RUN apk update && apk add --no-cache git
#RUN go get ./...

#COPY dependencies /go/src
RUN go mod init
RUN go get github.com/gin-gonic/gin
RUN go get github.com/dongri/phonenumber@latest

COPY src /go/src/go-docker-dev.to/src
COPY templates /go/src/go-docker-dev.to/templates

RUN go build go-docker-dev.to/src/app

EXPOSE $PORT

ENTRYPOINT ["./app"]
