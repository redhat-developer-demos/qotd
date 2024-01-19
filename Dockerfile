FROM docker.io/golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod init qotd
#RUN go install github.com/gorilla/mux@latest
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go build -o qotd .
EXPOSE 10000
CMD ["/app/qotd"]
