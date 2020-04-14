
all: clean build

clean:
	rm -rf bin

fmt:
	gofmt -w pkg/* main.go

vendor:
	go mod vendor

build:
	go build -o bin/server cmd/server.go
	go build -o bin/client cmd/client.go

proto:
	protoc -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway\
		--go_out=plugins=grpc:. pb/*.proto
	protoc -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway\
 		--grpc-gateway_out=logtostderr=true:. pb/*.proto
	protoc -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway\
		--swagger_out=logtostderr=true:. pb/*.proto

create-user:
	curl -X PUT -H "Content-Type: application/json" -d '{"name":"stijn"}'  localhost:8090/api/users

list-users:
	curl localhost:8090/api/users
