package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
	"time"
)

var DBConnection *gorm.DB
var MigrationPath = "file://migrations"

func InitDb() {
	conn, err := gorm.Open("postgres", "host=localhost port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	DBConnection = conn
	if err != nil {
		panic(err)
	}
}

func UpMigrate() {
	driver, err := postgres.WithInstance(DBConnection.DB(), &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		MigrationPath,
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}

func DownMigrate() {
	driver, err := postgres.WithInstance(DBConnection.DB(), &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		MigrationPath,
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	err = m.Down()
	if err != nil {
		panic(err)
	}
}

type Base struct {
	Id        string    `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	if base.Id == "" {
		id := uuid.NewV4().String()
		return scope.SetColumn("ID", id)
	}
	return nil
}
