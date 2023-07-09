FROM golang:1.20

ARG API_PORT

WORKDIR /app

COPY go.mod go.sum ./

COPY .env ./

RUN go mod download

COPY ./*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app

EXPOSE $API_PORT

CMD ["/bin/app"]