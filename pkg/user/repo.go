package user

import "github.com/stijndehaes/gin-gonic-gorm/pkg/db"

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
