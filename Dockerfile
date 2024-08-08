FROM golang:1.22-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM scratch

COPY --from=build /app/main /app

EXPOSE 8080

CMD ["/app/main"]
