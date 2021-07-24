# docker build -t spotifyalpine .
# docker run -p 8080:8080 -tid spotifyalpine

# heroku container:push web -a go-spotify-data
# heroku container:release web -a go-spotify-data

#########

## method 1 -> 246mb
# https://www.youtube.com/watch?v=FQS1p88Q0q8
FROM golang:alpine

WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app .

FROM python:slim

RUN pip install --no-cache-dir numpy pandas
COPY . .
COPY --from=0 /bin/app /bin/app

ENV IS_DOCKER=true
ENTRYPOINT ["/bin/app", "--port", "${PORT}"]
# EXPOSE 8080

#########

## method 2 -> 542mb
# FROM golang:alpine

# ADD . /go/src/gotify
# RUN apk add py3-numpy py3-pandas

# WORKDIR /go/src/gotify
# RUN go mod init gotify
# COPY . /go/src/gotify
# RUN go get github.com/gin-gonic/gin
# RUN go build -o /gotify

# ENV IS_DOCKER=true
# ENTRYPOINT /gotify --port $PORT
# # EXPOSE 8080

#########

## method 3 -> 1.25gb
# FROM golang:latest

# ADD . /go/src/gotify
# RUN apt-get update
# RUN apt-get install -y python3-pip
# RUN pip3 install 'numpy' && pip3 install 'pandas' && pip3 install 'pytz'

# WORKDIR /go/src/gotify
# RUN go mod init gotify
# COPY . /go/src/gotify
# RUN go get github.com/gin-gonic/gin
# RUN go build -o /gotify

# ENV IS_DOCKER=true
# ENTRYPOINT /gotify --port 8080 --host="0.0.0.0"
# EXPOSE 8080

#########