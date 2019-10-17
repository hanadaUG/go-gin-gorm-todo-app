package models

import (
	"github.com/hanadaUG/go-gin-gorm-todo-app/enum"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Text   string
	Status enum.State
}
