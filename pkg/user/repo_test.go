package user

import (
	"github.com/stijndehaes/gin-gonic-gorm/pkg/db"
	"github.com/stijndehaes/gin-gonic-gorm/pkg/house"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	db.InitDb()
	db.UpMigrate()
}

func shutdown() {
	db.DownMigrate()
}

func init() {
	db.MigrationPath = "file://../../migrations"
}

func TestListUsers(t *testing.T) {
	users, err := ListUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 0)
}

func TestInsertUser(t *testing.T) {
	user := User{Name: "Stijn"}
	err := InsertUser(&user)
	assert.NoError(t, err)
	assert.NotEqual(t, "", user.Id)
}

func TestUserAndHouses(t *testing.T) {
	user := User{Name: "Stijn"}
	err := InsertUser(&user)
	assert.NoError(t, err)
	h := house.House{
		Location: "Blanden",
		OwnerId:  user.Id,
	}
	err = house.InsertHouse(&h)
	assert.NoError(t, err)
	uh, err := UserAndHouses(user.Id)
	assert.NoError(t, err)
	assert.Len(t, *uh, 1)
	assert.Equal(t, []UserAndHouse{{
		UserId:        user.Id,
		UserName:      user.Name,
		HouseId:       h.Id,
		HouseLocation: h.Location,
	}}, *uh)
}
