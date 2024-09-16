package repoi

import (
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/models"
)

type TodoRepoI interface {
	CreateTodo(todo *models.NewTodo) error
	GetTodo(id *models.GetByID) (*models.Todo, error)
	GetTodos(req *models.Gets) (*[]models.Todo, error)
	UpdatedTodo(todo *models.Todo) error
	DeleteTodo(id *models.DeleteByID) (string, error)
}
