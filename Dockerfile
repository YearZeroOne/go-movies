FROM golang:1.22.0

WORKDIR /app

RUN go install github.com/air-verse/air@latest


COPY go.mod go.sum ./
RUN go mod download

RUN go mod tidy

CMD ["air", "-c", ".air.toml"]
