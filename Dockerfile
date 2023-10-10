FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /notification-receiver

EXPOSE 8888

# Run
CMD ["/notification-receiver"]