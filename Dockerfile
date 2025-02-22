FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN GOOS=linux go build -o /docker-yourquote

VOLUME  /quotes.sqlite

EXPOSE 54848

CMD ["/docker-yourquote"]

