FROM golang:1.15

RUN go get -u github.com/Sirupsen/logrus
RUN go get -v -u github.com/olivere/elastic
RUN go get -u -v github.com/go-zoo/bone

RUN mkdir -p /go/src/github.com/gunjan01/data_pipeline

ADD . /go/src/github.com/gunjan01/data_pipeline/
WORKDIR /go/src/github.com/gunjan01/data_pipeline/

RUN go build -o clarisights ./source/cmd...

CMD ["/go/src/github.com/gunjan01/data_pipeline/clarisights"]
