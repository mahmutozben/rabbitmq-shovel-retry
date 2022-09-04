FROM golang:1.14-alpine
RUN apk add --no-cache git
# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/mahmutozben/rabbitmq-shovel-retry
COPY . .
# Build the Go app
RUN go build -o /app/rabbitmq-shovel-retry .

EXPOSE 5000
ENTRYPOINT ["/app/rabbitmq-shovel-retry"]