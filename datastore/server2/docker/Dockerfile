FROM golang:1.11.5-alpine3.8
RUN apk add git
RUN go get "cloud.google.com/go/datastore" && \
    go get "github.com/gorilla/mux"
RUN mkdir /go/src/server
