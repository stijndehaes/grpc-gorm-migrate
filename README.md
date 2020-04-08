# Example project combining grpc with gorm and migrate

This project exposes both grpc as a rest api based on that grpc services using github.com/grpc-ecosystem/grpc-gateway.


##Running the project

To run this project first start the postgres server:

```bash
docker-compose up -d
```

Then download libraries and build the binaries:

```bash
make vendor
make build
```

Then you can start the server:

```bash
bin/server
```

In another terminal you can use the client:

```bash
bin/client user create -n "Stijn"
```

## Building the proto files

To build the proto files you should checkout grpc-gateway in your go path

```bash
cd $GOPATH/src/github.com/grpc-ecosystem
git clone https://github.com/grpc-ecosystem/grpc-gateway.git
```
This is because you need the proto annotations in that project, these are ignored
by go 1.14.x if you check them out in the vendor folder.

Additionally, you need to get the following plugins

```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

You also need to install [the protoc tool](http://google.github.io/proto-lens/installing-protoc.html).

After that you can run:

```bash
make proto
```
