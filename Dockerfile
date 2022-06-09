FROM golang:1.17 AS API
WORKDIR /sv-sso-gin
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o http_server ./main.go && mkdir run && mv ./http_server ./config.yaml run
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o rpc_server ./rpc.go && mv ./rpc_server run

FROM alpine:latest
ENV GIN_MODE=release

COPY --from=API /sv-sso-gin/run ./
CMD ["sh", "-c", "./http_server & \n ./rpc_server"]
