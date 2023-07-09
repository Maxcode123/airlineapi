FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum .env ./

RUN go mod download

COPY ./*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app

CMD ["/bin/app"]