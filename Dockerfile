FROM golang:1.20-alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o ./bin/app ./main.go 


FROM node:16-alpine
COPY --from=build-backend /build/main ./app/main
ADD ./frontend ./app/frontend

WORKDIR /app/frontend
RUN npm install

FROM alpine:3.14

# Install CA certificates to enable SSL/TLS certificate validation
RUN apk --no-cache add ca-certificates

WORKDIR /app

ENV ENV=dev

COPY --from=build /app/.env .
COPY --from=build /app/frontend /frontend
COPY --from=build /app/bin/app ./app

EXPOSE 8080

ENTRYPOINT ["./app"]