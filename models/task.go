package models

import (
	"github.com/hanadaUG/go-gin-gorm-todo-app/enum"
	"time"
)

type Task struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Text      string
	Status    enum.State
}
