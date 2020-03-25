FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get github.com/gorilla/mux
RUN go build -o qotd .
EXPOSE 8080
CMD ["/app/qotd"]