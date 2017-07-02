FROM golang:latest

WORKDIR /go/src/imageBoard
COPY . .

RUN go-wrapper download
CMD ["go-wrapper", "run"]
