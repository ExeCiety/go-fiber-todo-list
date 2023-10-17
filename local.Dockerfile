FROM golang:1.18.3-buster

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go install
RUN go install github.com/gofiber/cli/fiber@latest

EXPOSE 3001
CMD ["fiber", "dev"]
