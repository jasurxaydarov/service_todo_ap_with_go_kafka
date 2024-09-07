package storage

import (
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/storage/postgres"
	"github.com/jasurxaydarov/service_todo_ap_with_go_kafka/storage/repoi"
	"gorm.io/gorm"
)

type storage struct {
	todoRepoI repoi.TodoRepoI
}

type StorageI interface {
	GetTodoRepo() repoi.TodoRepoI
}

func NewStorage(db *gorm.DB) StorageI {

	return &storage{
		todoRepoI: postgres.NewTodoRepo(db),
	}
}

func (s *storage) GetTodoRepo() repoi.TodoRepoI {

	return s.todoRepoI
}
