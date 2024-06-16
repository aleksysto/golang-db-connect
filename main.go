package main

import (
	"github.com/aleksysto/golang-db-connect.git/db"
   // "github.com/aleksysto/golang-db-connect.git/packer"
)

func main() {
    driver := db.NewDriver()
    driver.OpenConnection("bolt+://localhost:7687")
	
}
