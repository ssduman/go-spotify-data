# https://stackoverflow.com/a/52474413
# docker build -t spotifyalpine .
# docker run -p 8080:8080 -tid spotifyalpine
FROM golang:latest

ADD . /go/src/gotify

RUN apt-get update
RUN apt-get install -y python3-pip
RUN pip3 install 'numpy' && pip3 install 'pandas' && pip3 install 'pytz'

WORKDIR /go/src/gotify
RUN go mod init gotify
COPY . /go/src/gotify
RUN go get github.com/gin-gonic/gin
RUN go build -o /gotify

ENTRYPOINT /gotify --port 8080 --host="0.0.0.0"

EXPOSE 8080

# FROM golang:alpine

# ADD . /go/src/gotify

# RUN apk update
# RUN apk add --no-cache musl-dev linux-headers g++ libblas-dev liblapack-dev
# RUN apk add --update make cmake gcc g++ gfortran
# RUN apk add --update --no-cache python3 py3-pip python3-dev && ln -sf python3 /usr/bin/python
# RUN python3 -m ensurepip
# RUN pip3 install --no-cache --upgrade pip setuptools cython
# RUN pip3 install 'numpy' && pip3 install 'pandas' && pip3 install 'pytz'

# WORKDIR /go/src/gotify
# RUN go mod init gotify
# COPY . /go/src/gotify
# RUN go get github.com/gin-gonic/gin
# RUN go build -o /gotify

# ENTRYPOINT /gotify --port 8080 --host="0.0.0.0"

# EXPOSE 8080