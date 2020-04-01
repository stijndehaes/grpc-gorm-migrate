
all: clean build

clean:
	rm -rf $(TARGET)

fmt:
	gofmt -w pkg/* main.go

vendor:
	go mod vendor

build:
	go build main.go

proto:
	protoc -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
		--go_out=plugins=grpc:. pb/*.proto
	protoc -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
 		--grpc-gateway_out=logtostderr=true:. pb/*.proto
	protoc -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
		--swagger_out=logtostderr=true:. pb/*.proto

create-user:
	curl -X PUT -H "Content-Type: application/json" -d '{"name":"stijn"}'  localhost:8090/api/users

list-users:
	curl localhost:8090/api/users
