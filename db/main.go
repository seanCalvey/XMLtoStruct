package db

import (
	"github.com/go-pg/pg"
	"log"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "password",
		Addr:     "localhost:5432",
		Database: "fobar",
	}
	db := pg.Connect(opts)
	if db == nil {
		log.Panic("failed to connect to db")
	}
	log.Printf("connected to Db\n")

	err := CreateFooTable(db)
	if err != nil {
		log.Panic("Table not Created")
	}

	err = CreateBarTable(db)
	if err != nil {
		log.Panic("Table not Created")
	}

	return db
}
