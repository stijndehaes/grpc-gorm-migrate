package house

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stijndehaes/gin-gonic-gorm/pkg/db"
	"github.com/stijndehaes/gin-gonic-gorm/pkg/user"
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

func TestListHouses(t *testing.T) {
	houses, err := ListHouses()
	assert.NoError(t, err)
	assert.Len(t, houses, 0)
}

func TestInsertHouse(t *testing.T) {
	u := user.User{
		Name: "Stijn",
	}
	err := user.InsertUser(&u)
	assert.NoError(t, err)
	house := House{
		Location: "Blanden",
		OwnerId:  u.ID,
	}
	err = InsertHouse(&house)
	assert.NoError(t, err)
	assert.NotEqual(t, "", house.ID)
}

func TestInsertHouse_FailsWithoutOwner(t *testing.T) {
	house := House{
		Location: "Blanden",
		OwnerId:  uuid.NewV4().String(),
	}
	err := InsertHouse(&house)
	assert.Error(t, err)
}
