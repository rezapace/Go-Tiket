FROM golang:1.21-alpine

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN migrate -database "postgres://postgres:mysecretpassword@localhost:5432/Tiketing?sslmode=disable" -path db/migration-golang up

RUN go build -o app .

CMD ["/app/app"]