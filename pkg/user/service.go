package user

import (
	"context"
	"github.com/jinzhu/copier"
	pb "github.com/stijndehaes/grpc-gorm-migrate/pb"
)

type Service struct {
}

func (g *Service) UserWithHouses(c context.Context, r *pb.UserHousesRequest) (*pb.UserHousesResponse, error) {
	uh, err := UserAndHouses(r.Id)
	if err != nil {
		return nil, err
	}
	var uhr []*pb.UserHouse
	err = copier.Copy(&uhr, &uh)
	if err != nil {
		return nil, err
	}
	return &pb.UserHousesResponse{
		UserHouses: uhr,
	}, err
}

func (g *Service) GetUsers(ctx context.Context, r *pb.UsersRequest) (*pb.UsersResponse, error) {
	users, err := ListUsers()
	if err != nil {
		return nil, err
	}
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:   user.Id,
			Name: user.Name,
		})
	}
	return &pb.UsersResponse{
		Users: pbUsers,
	}, nil
}

// SayHello says hello
func (g *Service) CreatUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	createdUser := User{
		Name: in.GetName(),
	}
	err := InsertUser(&createdUser)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:   createdUser.Id,
		Name: createdUser.Name,
	}, nil
}
