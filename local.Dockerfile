FROM golang:1.18.3-buster

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go install
RUN go install github.com/cosmtrek/air@latest
#RUN air init

EXPOSE 3001
CMD ["air", "-c", ".air.docker.toml"]
