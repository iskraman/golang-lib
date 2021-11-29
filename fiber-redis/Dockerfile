FROM golang:1.17.3

RUN apt update && apt upgrade -y

RUN mkdir -p /app/src
COPY * /app/src/
WORKDIR /app/src

RUN go mod tidy

CMD ["go", "run", "main.go"]
