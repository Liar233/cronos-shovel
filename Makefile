.PHONY: build run stop test lint protogen

build:
	docker network rm -f global-network
	docker network rm -f cronos-shovel-network
	docker network create global-network
	docker network create cronos-shovel-network
	docker-compose -f ./build/docker-compose.yml build --no-cache $(c)
run:
	docker-compose -f ./build/docker-compose.yml up -d $(c)
stop:
	docker-compose -f ./build/docker-compose.yml down $(c)
test:
	docker-compose -f ./build/docker-compose.yml exec cronos-shovel-server go test -race ./... $(c)
lint:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v1.54.2 golangci-lint run -v $(c)
generate:
	docker-compose -f build/docker-compose.yml exec cronos-shovel-server protoc --go_out=./internal/ --go_opt=paths=source_relative --go-grpc_out=./internal/ --go-grpc_opt=paths=source_relative ./pkg/*.proto$(c)
