FROM golang:1.12

WORKDIR $GOPATH/src/github.com/pqnguyen

COPY . .

ENV GO111MODULE=on
ENV GIT_TERMINAL_PROMPT=1

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["api"]