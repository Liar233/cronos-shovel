FROM golang:1.21.1 AS build-server

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y protobuf-compiler \
    && go get google.golang.org/protobuf/cmd/protoc-gen-go \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go \
    && go get google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    && export PATH="$PATH:$(go env GOPATH)/bin" \
 && rm -rf /var/lib/apt/lists/*

RUN go mod download \
    && go mod verify \
    && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /output/server /app/cmd/server/ \
    && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /output/client /app/cmd/cli/

RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint

RUN chmod 0755 /output/server
RUN chmod +x /output/server
RUN chmod 0755 /output/client
RUN chmod +x /output/client


ENTRYPOINT ["/output/server", "/app/configs/config.example.yml"]
