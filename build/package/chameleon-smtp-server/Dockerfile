FROM golang:1.16-alpine
ADD . /go/src/github.com/YuKitsune/Chameleon
WORKDIR /go/src/github.com/YuKitsune/Chameleon
RUN go build -o bin/chameleon-smtp-server cmd/chameleon-smtp-server/main.go
CMD ["./bin/chameleon-smtp-server"]