FROM golang:latest

COPY . /src/app
WORKDIR /src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --polling --log-prefix=false --build="go build main" --command="./main" --directory="./"