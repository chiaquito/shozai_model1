FROM golang:1.23

WORKDIR /home/shozai_model

# Port 1323 for application, 2345 for Delve
EXPOSE 8080 2345

# install air, delve 
RUN go install github.com/cosmtrek/air@v1.42.0 && \ 
    go install github.com/go-delve/delve/cmd/dlv@latest 

COPY ["./go.mod", "./go.sum", "./main.go", "/home/shozai_model/"]
RUN go mod download

ENTRYPOINT [ "air", "-c", ".air.toml"]