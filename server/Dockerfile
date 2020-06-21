FROM golang:1.14.2 AS build
WORKDIR /app
COPY . .
RUN go mod download && CGO_ENABLED=0 go build -o build/server .

FROM scratch
WORKDIR /app
COPY --from=build /app/build/server server

EXPOSE 8000
ENTRYPOINT ["/app/server"]