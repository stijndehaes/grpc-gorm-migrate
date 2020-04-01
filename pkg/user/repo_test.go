package user

import (
	"github.com/stijndehaes/gin-gonic-gorm/pkg/db"
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
	db.DBConnection.AutoMigrate(&User{})
	users, err := ListUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 0)
}

func TestInsertUser(t *testing.T) {
	db.DBConnection.AutoMigrate(&User{})
	user := User{Name: "Stijn"}
	err := InsertUser(&user)
	assert.NoError(t, err)
	assert.NotEqual(t, "", user.ID)
}
