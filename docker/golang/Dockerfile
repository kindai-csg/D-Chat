FROM golang:1.14.3-buster

ENV GOPATH /.go

RUN apt-get -y update && \
    apt-get -y install git

RUN go get -v github.com/gin-gonic/gin

# development 
RUN go get -v github.com/golang/mock/gomock
RUN go get -v golang.org/x/mod/modfile
RUN go install github.com/golang/mock/mockgen

CMD ["go", "run", "main.go"]