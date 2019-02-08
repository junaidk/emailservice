# build stage
FROM golang:1.11.3  AS build-env

# Set our workdir to our current service in the gopath
WORKDIR /go/src/emailservice/
# Copy the current code into our workdir
COPY . .
ENV GOPATH /go/
RUN go build -o emailservice main/main.go

# final stage
FROM ubuntu:bionic

WORKDIR /app

COPY --from=build-env /go/src/emailservice/template/ /app/template/

COPY --from=build-env /go/src/emailservice/emailservice /app/


CMD ./emailservice
