package house

import (
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/db"
)

type House struct {
	db.Base
	Location string
	OwnerId  string
}
