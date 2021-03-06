package user

import (
	"database/sql"
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/server/db"
)

type User struct {
	db.Base
	Name string        `json:"name"`
	Age  sql.NullInt64 `json:"age"`
}
