package user

import (
	"database/sql"
	"github.com/stijndehaes/gin-gonic-gorm/pkg/db"
)

type User struct {
	db.Base
	Name string        `json:"name"`
	Age  sql.NullInt64 `json:"age"`
}
