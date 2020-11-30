FROM golang:1.15-alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download

WORKDIR /app/cmd
RUN go build -o /bin/main
WORKDIR /app
CMD ["/bin/main"]