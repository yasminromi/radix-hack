FROM golang AS build-env

ADD . /go/src/github.com/yasminromi/radix-hack

RUN mkdir /app && \
    cd /go/src/github.com/yasminromi/radix-hack && \
    CGO_ENABLED=0 GOOS=linux go build -v -ldflags '-s' -a -installsuffix cgo -o service-go . && \
    mv go/src/github.com/yasminromi/radix-hack/service-go /app/ && \
    rm -rf /go/src/


FROM alpine

RUN apk update && \
   apk add ca-certificates && \
   update-ca-certificates && \
   apk add --no-cache tzdata && \
   rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=build-env /app/service-go /app

ENV TZ=America/Sao_Paulo

RUN echo $TZ > /etc/timezone

ENV GO_PORT 80

EXPOSE 80

ENTRYPOINT ./service-go
