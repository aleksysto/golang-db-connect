package main

import (
	"github.com/aleksysto/golang-db-connect.git/db"
)

func main() {
	d := db.NewDriver()
	d.OpenConnection("bolt://localhost:7687")
}
