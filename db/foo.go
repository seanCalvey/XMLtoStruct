package db

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
)

type FooResult struct {
	Version string `xml:"version" sql:"version"`
	Time    string `xml:"time" sql:"time,type:real"`
	Bar     *Bar   `xml:"bar" sql:"-"`
}

type Bar struct {
	Name  string `xml:"name"`
	Order string `xml:"order"`
}

func CreateFooTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&FooResult{}, opts)

	if err != nil {
		log.Printf("Can't create table: %v", err)
		return err
	}
	log.Println("Table Created")
	return nil
}

func CreateBarTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&Bar{}, opts)

	if err != nil {
		log.Printf("Can't create table: %v", err)
		return err
	}
	log.Println("Table Created")
	return nil
}

func (f *FooResult) Insert(db *pg.DB) error {
	err := db.Insert(f)
	if err != nil {
		log.Printf("Can't insert foo: %v", err)
		return err
	}
	err = db.Insert(f.Bar)
	log.Println("Foo Inserted")
	return nil
}
