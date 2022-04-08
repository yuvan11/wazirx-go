FROM golang:1.16-alpine
WORKDIR /go/src/app
RUN echo $GOPATH
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
#WORKDIR $GOPATH/src
#RUN echo $GOPATH
CMD ["go", "run", "./cmd/main.go"]
#CMD ["app"]
#EXPOSE 8081

