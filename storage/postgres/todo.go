package postgres

import (
	"fmt"

	"github.com/google/uuid"
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

func (t *todoRepo) CreateTodo(todo *models.NewTodo) error {

	req := models.Todo{
		Id:   uuid.New(),
		Task: todo.Task,
	}

	res := t.db.Create(&req)

	return res.Error
}

func (t *todoRepo) GetTodo(id *models.GetByID) (*models.Todo, error) {

	var todo models.Todo

	res := t.db.First(&todo, "id = ?", id)

	if res.Error != nil {

		fmt.Println("erron get todo", res.Error.Error())

		return nil, res.Error
	}

	return &todo, nil
}

func (t *todoRepo) GetTodos(req *models.Gets) (*[]models.Todo, error) {

	var todos []models.Todo

	res := t.db.Limit(req.Limit).Offset(req.Offset).Find(&todos)

	if res.Error != nil {
		fmt.Println("erron get todos", res.Error.Error())

		return nil, res.Error
	}

	return &todos, nil
}

func (t *todoRepo) DeleteTodo(id *models.DeleteByID) (string, error) {


	// DeleteTodo deletes a Todo by its UUID.
		// Find the Todo item by ID
		var todo models.Todo
		res := t.db.First(&todo, "id = ?", id)
		if res.Error != nil {
			if res.Error == gorm.ErrRecordNotFound {
				return "", fmt.Errorf("todo not found")
			}
			return "", res.Error
		}
	
		// Delete the Todo item
		res = t.db.Delete(&todo)
		if res.Error != nil {
			fmt.Println("error deleting todo:", res.Error.Error())
			return "", res.Error
		}
	
		return "successfully deleted", nil
	}


func (t *todoRepo) UpdatedTodo(todo *models.Todo) error {

	res := t.db.Save(&todo)

	if res.Error != nil {

		fmt.Println("err on update todo", res.Error.Error())
	}
	return nil
}
