package connection

import (
	"fmt"
	pb "github.com/stijndehaes/grpc-gorm-migrate/pb"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var HouseClient pb.HouseServiceClient
var UserClient pb.UserServiceClient

func InitClients() {
	fmt.Sprintln("Conection to the server")
	connection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	conn = connection
	HouseClient = pb.NewHouseServiceClient(conn)
	UserClient = pb.NewUserServiceClient(conn)
}

func CloseConnection() {
	_ = conn.Close()
}
