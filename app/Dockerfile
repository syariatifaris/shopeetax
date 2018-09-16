FROM golang

COPY ./ /go/src/github.com/syariatifaris/shopeetax/app/
WORKDIR /go/src/github.com/syariatifaris/shopeetax/app/

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure -v