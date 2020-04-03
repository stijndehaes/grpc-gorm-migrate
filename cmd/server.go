package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/server"
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/server/db"
)

func main() {
	db.InitDb()
	db.UpMigrate()
	log.SetLevel(log.DebugLevel)
	g := server.New()
	g.Start()
	g.WaitStop()
}
