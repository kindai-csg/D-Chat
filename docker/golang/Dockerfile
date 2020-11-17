FROM golang:1.14.3-buster

ENV GOPATH /.go

RUN apt-get -y update && \
    apt-get -y install git

# development 
RUN go get -v github.com/golang/mock/mockgen
RUN go install -v github.com/golang/mock/mockgen

CMD ["go", "run", "main.go"]
