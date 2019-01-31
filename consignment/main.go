package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ConsignmentReposityory --
type ConsignmentRepository interface{}

type Repository struct {
	DB *gorm.DB
}

var (
	// db
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbName = os.Getenv("DB_NAME")
	dbPass = os.Getenv("DB_PASS")
	dbSsl  = os.Getenv("DB_SSL")
)

func main() {
	// db, err := db.GetDB(dbHost, dbPort, dbUser, dbPass, dbName, dbSsl)
	// if err != nil {
	// 	log.Fatalf("cannot start db: %v", err)
	// }
	// repo := &Repository{DB: db}

	// New Service
	// s := micro.NewService(
	// 	micro.Name("go.micro.srv.consignment"),
	// 	micro.Version("latest"),
	// )

	// Initialise service
	// s.Init()

	// register to auto-genterated code
	// pb.RegisterShippingServiceHandler(s.Server(), &service{repo})

	// Run service
	// if err := s.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
