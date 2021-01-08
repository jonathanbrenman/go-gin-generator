FROM golang:1.14-alpine
RUN apk add git gcc g++
ENV CGO_ENABLED 1
ENV GOPATH /go
ENV CC gcc


ADD . /go/src/go-app
RUN cd /go/src/go-app && go mod tidy
WORKDIR /go/src/go-app/cmd/api

RUN go get -u github.com/cosmtrek/air
ENTRYPOINT  air run main.go
EXPOSE 8080
