package house

import (
	"github.com/stijndehaes/gin-gonic-gorm/pkg/db"
)

type House struct {
	db.Base
	Location string
	OwnerId  string
}
