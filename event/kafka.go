package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/dto"
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/models"
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/storage"
	"github.com/segmentio/kafka-go"
)

const (
	KafkaBroker = "localhost:9092"
	GroupID     = "crud-service-group"
)

func consumeMassages(topic string, storage storage.StorageI) {
	reader := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: []string{KafkaBroker},
			Topic:   topic,
		},
	)

	for {
		msg, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Println("err on msg,err:=reader.ReadMessage(context.Background())", err)
			continue
		}

		switch topic{
		case "create-todo":
			var todo dto.CreateTodoDto

			log.Println(string(msg.Value))
			if err:=json.Unmarshal(msg.Value,&todo);err!=nil{
				log.Println("could not unmarshal message:", err)
				continue
			}

			if err:=storage.GetTodoRepo().CreateTodo(&todo);err!=nil{
				log.Println(err)
				continue
			}

		case "update-todo":
			var todo models.TodoModel

			log.Println(string(msg.Value))
			if err:=json.Unmarshal(msg.Value,&todo);err!=nil{
				log.Println("could not unmarshal message:", err)
				continue
			}

			if err:=storage.GetTodoRepo().UpdatedTodo(&todo);err!=nil{
				log.Println(err)
				continue
			}
		}
	}
}
