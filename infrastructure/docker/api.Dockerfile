FROM golang:1.25-alpine AS build

WORKDIR /app

COPY apps/api/go.mod apps/api/go.sum* ./
RUN go mod download

COPY apps/api ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server

FROM alpine:3.21

WORKDIR /app
COPY --from=build /server /server

EXPOSE 8080

CMD ["/server"]
