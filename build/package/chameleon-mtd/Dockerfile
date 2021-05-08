FROM golang:1.16-alpine

# Versioning information
ARG GIT_COMMIT
ARG GIT_BRANCH=main
ARG GIT_DIRTY='false'
ARG VERSION
LABEL branch=$GIT_BRANCH \
    commit=$GIT_COMMIT \
    dirty=$GIT_DIRTY \
    version=$VERSION

ADD . /go/src/github.com/YuKitsune/Chameleon
WORKDIR /go/src/github.com/YuKitsune/Chameleon
RUN go build -o bin/chameleon-mtd -ldflags "-X 'github.com/yukitsune/chameleon.Version=$VERSION'" cmd/chameleon-mtd/main.go
CMD ["./bin/chameleon-mtd", "serve"]