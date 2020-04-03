package house

import (
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/server/db"
)

type House struct {
	db.Base
	Location string
	OwnerId  string
}
