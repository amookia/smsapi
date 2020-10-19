FROM golang:1.14

WORKDIR /go/src/app
COPY /gowebsite .

RUN go mod download
RUN go install main.go

CMD ["main"]