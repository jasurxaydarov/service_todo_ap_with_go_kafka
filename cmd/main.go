package main

import (
	"log"

	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/event"
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/storage"
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/storage/postgres"
)

func main() {

	db, err := postgres.ConnToDb()

	if err != nil {

		log.Println("err on db,err:=postgres.ConnToDb()", err)
		return
	}
	log.Println(db)

	// postgres.AoutoMigrate(db)
	storage := storage.NewStorage(db)

	event := event.NewEvent(storage)

	event.Run()

}
