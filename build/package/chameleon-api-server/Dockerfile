FROM golang:1.16-alpine
ADD . /go/src/github.com/YuKitsune/Chameleon
WORKDIR /go/src/github.com/YuKitsune/Chameleon
RUN go build -o chameleon-api-server cmd/chameleon-api-server/main.go
CMD ["./chameleon-api-server"]