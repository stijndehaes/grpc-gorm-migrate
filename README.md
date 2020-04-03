# Example project combining grpc with gorm and migrate

This project exposes both grpc as a rest api based on that grpc services using github.com/grpc-ecosystem/grpc-gateway.

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
