FROM golang:1.22.5

COPY . .

RUN go build -o server .

EXPOSE 80

CMD ["./server"]