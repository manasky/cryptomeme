FROM golang:1.16

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -a -ldflags "-w -s" -o ./bin/crypromeme main.go

EXPOSE 8081
ENTRYPOINT ["./bin/cryptomeme"]