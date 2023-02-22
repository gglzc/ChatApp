#Build Image
FROM        golang:1.18.4-alpine3.16 AS builder 
WORKDIR     /app
COPY        . .
COPY        .env .
RUN         go build -o main main.go
ENTRYPOINT  ["./app"]

#Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY config/config.yml /app/config/config.yml
 
EXPOSE 8085
CMD ["/app/main"]
