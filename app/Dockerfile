FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o server .

EXPOSE 10102

CMD ["./server"]

