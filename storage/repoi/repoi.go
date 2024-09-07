package repoi

import (
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/dto"
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/models"
)

type TodoRepoI interface {
	CreateTodo(todo *dto.CreateTodoDto) error
	UpdatedTodo(todo *models.TodoModel) error
}
