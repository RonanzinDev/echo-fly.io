FROM golang:1.20-alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o ./bin/app ./cmd/main/main.go 

FROM alpine:3.14

# Install CA certificates to enable SSL/TLS certificate validation
RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=build /app/.env .

COPY --from=build /app/view /view

COPY --from=build /app/bin/app ./app

EXPOSE 8080

ENTRYPOINT ["./app"]