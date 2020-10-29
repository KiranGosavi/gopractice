FROM golang:latest

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o demo1 .

EXPOSE 5005

CMD ["/app/demo1"]