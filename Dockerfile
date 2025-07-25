FROM golang:1.24.5

WORKDIR /GoAssignment

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
