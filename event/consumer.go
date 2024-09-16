package event

import (
	"log"

	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/storage"
)

type event struct {
	str storage.StorageI
}

func NewEvent(str storage.StorageI) *event {

	return &event{str: str}
}

func (e *event) createTodo() {
	topic := "create-todo"
	log.Println(topic, " :topic is runing")

	consumeMassages(topic, e.str)

}

func (e *event) getTodo() {
	topic := "get-todo"
	log.Println(topic, " :topic is runing")

	consumeMassages(topic, e.str)

}

func (e *event) getTodos() {
	topic := "get-todos"
	log.Println(topic, " :topic is runing")

	consumeMassages(topic, e.str)

}

func (e *event) updateTodo() {
	topic := "update-todo"
	log.Println(topic, " :topic is runing")

	consumeMassages(topic, e.str)

}

func (e *event) deleteTodo() {
	topic := "delete-todo"
	log.Println(topic, " :topic is runing")

	consumeMassages(topic, e.str)

}

func (e *event) Run() {
	go e.createTodo()
	go e.updateTodo()
	go e.deleteTodo()
	select {}
}
