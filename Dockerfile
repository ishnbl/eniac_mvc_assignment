FROM golang:1.26.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


EXPOSE 8000

CMD ["go", "run", "main.go"]