package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stijndehaes/gin-gonic-gorm/pkg/db"
	"github.com/stijndehaes/gin-gonic-gorm/pkg/server"
)

//func main() {
//	parser := argparse.NewParser("datafy-ui", "Shows you the datafy UI")
//	logLevel := parser.String("", "log-level", &argparse.Options{Required: false, Help: "logging level to use", Default: "info"})
//	initLogging(*logLevel)
//	db.InitDb()
//	defer db.DBConnection.Close()
//	db.DBConnection.AutoMigrate(&user.User{})
//}
//
//func initLogging(levelString string) {
//	log.SetFormatter(&log.JSONFormatter{})
//	level, err := log.ParseLevel(levelString)
//	if err != nil {
//		panic(err)
//	}
//	log.SetLevel(level)
//}

func main() {
	db.InitDb()
	db.UpMigrate()
	log.SetLevel(log.DebugLevel)
	g := server.New()
	g.Start()
	g.WaitStop()
}
