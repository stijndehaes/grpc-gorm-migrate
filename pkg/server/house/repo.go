package house

import "github.com/stijndehaes/grpc-gorm-migrate/pkg/server/db"

func ListHouses() ([]House, error) {
	var houses []House
	err := db.DBConnection.Find(&houses).Error
	if err != nil {
		return nil, err
	}
	return houses, nil
}

func InsertHouse(house *House) error {
	return db.DBConnection.Create(house).Error
}
