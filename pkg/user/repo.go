package user

import (
	"github.com/stijndehaes/gin-gonic-gorm/pkg/db"
)

func ListUsers() ([]User, error) {
	var users []User
	err := db.DBConnection.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func InsertUser(user *User) error {
	return db.DBConnection.Create(user).Error
}

type UserAndHouse struct {
	UserId        string
	UserName      string
	HouseId       string
	HouseLocation string
}

func UserAndHouses(userId string) (*[]UserAndHouse, error) {
	var uh []UserAndHouse
	err := db.DBConnection.Table("users").
		Select("users.id as user_id, users.name as user_name, houses.id as house_id, houses.location as house_location").
		Where("users.id = ?", userId).
		Joins("JOIN houses on houses.owner_id = users.id").
		Scan(&uh).Error
	if err != nil {
		return nil, err
	}
	return &uh, nil
}
