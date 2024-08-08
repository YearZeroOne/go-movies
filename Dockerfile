FROM golang:1.22.0

WORKDIR /app

RUN go install github.com/air-verse/air@latest


COPY go.mod go.sum ./
RUN go mod download

EXPOSE 3000

CMD ["air", "-c", ".air.toml"]
