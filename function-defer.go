package main

import (
	"errors"
	"fmt"
)

type DB struct {
	opened bool
}

func open() *DB {
	db := new(DB)
	db.opened = true
	fmt.Println("connections is open")
	return db
}

func save(db *DB) error {
	if !db.opened {
		return errors.New("connection is closed")
	}
	fmt.Println("data is saved")
	return nil
}

func closeDb(db *DB) {
	db.opened = false
	fmt.Println("connections is closed")
}

func main() {
	db := open()
	defer closeDb(db)
	err := save(db)
	if err != nil {
		println(err)
	}
}
