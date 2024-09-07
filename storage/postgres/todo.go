package postgres

import (
	"github.com/google/uuid"
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/dto"
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/models"
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/storage/repoi"
	"gorm.io/gorm"
)

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) repoi.TodoRepoI {

	return &todoRepo{db: db}
}

func (t *todoRepo) CreateTodo(todo *dto.CreateTodoDto) error {

	req := models.TodoModel{
		Id:   uuid.New(),
		Task: todo.Task,
	}

	tx := t.db.Create(&req)
	return tx.Error
}

func (t *todoRepo) UpdatedTodo(todo *models.TodoModel) error {

	t.db.Save(&todo)

	return nil
}
