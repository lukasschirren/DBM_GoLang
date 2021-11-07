package db

import (
	"log"
	"os"

	pg "github.com/go-pg/pg/v10"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User: "postgres",
		Addr: "localhost:5432",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}
	log.Printf("Connection to database successful.\n")
	// CreateProjectTable(db)
	// CreateInvestmentTable(db)
	// CreateStatusTable(db)
	// CreateStatusTypeTable(db)
	// CreateFromToTable(db)
	// CreateCountryTable(db)
	// CreateTechnologyTable(db)
	CreateTypeTable(db)

	// closeErr := db.Close()
	// if closeErr != nil {
	// 	log.Printf("Error while closing the connection, Reason: %v\n", closeErr)
	// 	os.Exit(100)
	// }
	// log.Printf("Connection closed successfully. \n")
	return db
}
