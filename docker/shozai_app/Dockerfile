FROM golang:1.23

WORKDIR /home/shozai_model

# RUN go install github.com/cosmtrek/air@v1.42.0 && \
#     go install github.com/go-delve/delve/cmd/dlv@latest

COPY ["./go.mod", "./go.sum", "./main.go", "/home/shozai_model/"]
RUN mkdir internal && go mod download

COPY ["./internal",  "./internal/"]