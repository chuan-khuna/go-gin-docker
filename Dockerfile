FROM golang:1.21.1-bullseye

WORKDIR /app
COPY . /app

RUN go install github.com/cosmtrek/air@latest

RUN go get -u github.com/gin-gonic/gin
RUN go mod download

CMD ["air", "-c", ".air.toml"]