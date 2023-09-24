FROM golang:latest

WORKDIR /server

COPY go.mod /server/
COPY go.sum /server/

RUN go mod download

COPY . .

RUN go build -o server

EXPOSE 8080

CMD ["./server"]