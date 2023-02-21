#Build Image
FROM        golang:1.18.4-alpine3.16 AS builder 
WORKDIR     /app
COPY        . .
RUN         go build -o main main.go
ENTRYPOINT  ["./app"]

#Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8085
CMD ["/app/main"]
