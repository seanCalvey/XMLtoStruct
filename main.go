package main

import (
	"XMLtoStruct/db"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {

	// Open our xmlFile
	xmlFile, err := os.Open("info.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened info.xml")

	// defer the closing of our xmlFile so that we can parse it later on
	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {

		}
	}(xmlFile)

	byteValue, _ := io.ReadAll(xmlFile)

	// we initialize our Users array
	var results db.FooResult
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'results' which we defined above
	err = xml.Unmarshal(byteValue, &results)
	if err != nil {
		fmt.Println(err)
	}

	dbPool := db.Connect()

	err = db.CreateFooTable(dbPool)
	if err != nil {
		fmt.Println(err)
	}

	err = results.Insert(dbPool)
	if err != nil {
		fmt.Println(err)
	}
}
