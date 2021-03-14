# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todo cmd/todo/main.go

# final stage
FROM scratch
COPY --from=builder /app/todo /app/
EXPOSE 8080
ENTRYPOINT ["/app/todo"]