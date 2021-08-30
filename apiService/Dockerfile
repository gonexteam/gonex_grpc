FROM golang:1.15 AS builder

# Set GOPATH to build Go app
#ENV GOPATH=/go
#
## Set apps source directory
#ENV SRC_DIR=${GOPATH}/src/gonex.net/
#
## Define current working directory
#WORKDIR ${SRC_DIR}
#
## Copy apps scource code to the image
#COPY . ${SRC_DIR}

WORKDIR /app
COPY apiService /app
COPY protos /app
COPY go.mod /app
COPY go.sum /app
COPY goBuild.sh /app

# Build App
RUN ./goBuild.sh

EXPOSE 50059
ENTRYPOINT [ "apiService" ]
