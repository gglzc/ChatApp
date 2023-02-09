FROM        golang:1.18.4
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN         go mod download && go build -o app
ENTRYPOINT  ["./app"]