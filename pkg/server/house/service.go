package house

import (
	"context"
	"github.com/jinzhu/copier"
	pb "github.com/stijndehaes/grpc-gorm-migrate/pb"
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/server/db"
)

type Service struct {
	pb.HouseServiceServer
}

func (s *Service) GetHouses(c context.Context, r *pb.HousesRequest) (*pb.HousesResponse, error) {
	houses, err := ListHouses()
	if err != nil {
		return nil, err
	}
	var pHouses []*pb.House
	err = copier.Copy(&pHouses, &houses)
	if err != nil {
		return nil, err
	}
	return &pb.HousesResponse{
		Houses: pHouses,
	}, nil
}

func (s *Service) CreateHouse(c context.Context, r *pb.CreateHouseRequest) (*pb.House, error) {
	house := House{}
	err := copier.Copy(&house, r)
	if err != nil {
		return nil, err
	}
	err = InsertHouse(&house)
	if err != nil {
		return nil, err
	}
	pbHouse := pb.House{}
	err = copier.Copy(&pbHouse, house)
	if err != nil {
		return nil, err
	}
	return &pbHouse, nil
}

func (s *Service) StreamHouses(r *pb.HousesRequest, rs pb.HouseService_StreamHousesServer) error {
	rows, err := db.DBConnection.Table("houses").Rows()
	if err != nil {
		return err
	}
	for rows.Next() {
		var h House
		err = db.DBConnection.ScanRows(rows, &h)
		if err != nil {
			return err
		}

		ph := pb.House{}
		err := copier.Copy(&ph, &h)
		if err != nil {
			return err
		}
		rs.Send(&ph)
	}
	return nil
}
