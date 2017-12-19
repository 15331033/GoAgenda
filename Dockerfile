FROM golang:1.8

COPY . $GOPATH/src/github.com/caijh23/GoAgenda
WORKDIR $GOPATH/src/github.com/caijh23/GoAgenda/cli

RUN go-wrapper download
RUN go build -o $GOPATH/bin/agendalocal .

WORKDIR $GOPATH/src/github.com/caijh23/GoAgenda/http-api

RUN go-wrapper download
RUN go build -o $GOPATH/bin/agendaserver .

CMD ["agendaserver"]

EXPOSE 8080