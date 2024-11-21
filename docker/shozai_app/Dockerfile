FROM golang:1.23 AS builder

# Port 1323 for application, 2345 for Delve
EXPOSE 8080 2345

WORKDIR /home/shozai_model

COPY ["./go.mod", "./go.sum", "./main.go", "/home/shozai_model/"]
RUN go mod download


# 開発環境用のステージ
FROM builder AS dev

# install air, delve 
RUN go install github.com/cosmtrek/air@v1.42.0 && \ 
    go install github.com/go-delve/delve/cmd/dlv@latest 

ENTRYPOINT [ "air", "-c", ".air.toml"]


# 本番環境用のステージ
FROM golang:1.23 AS prod
WORKDIR /home/shozai_model
COPY [".", "/home/shozai_model/"]
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o ./dist/main .
CMD ["./dist/main"]