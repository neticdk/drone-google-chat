FROM alpine:3.6 as alpine
RUN apk add -U --no-cache ca-certificates

##
## Build
##
FROM golang:1.18 AS build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...
RUN go build -o /go/bin/app

##
## Distribution image
##
FROM gcr.io/distroless/base-debian11
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
LABEL org.label-schema.version=latest
LABEL org.label-schema.vcs-url="https://github.com/josmo/drone-google-chat.git"
LABEL org.label-schema.name="Drone Google Chat"
LABEL org.label-schema.vendor="Josmo"
COPY --from=build /go/bin/app /
CMD ["/app"]
