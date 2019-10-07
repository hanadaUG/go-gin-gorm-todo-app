package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hanadaUG/go-gin-gorm-todo-app/models"
	"github.com/jinzhu/gorm"
	"net/http"
)

type TaskHandler struct {
	Db *gorm.DB
}

// 一覧表示
func (handler *TaskHandler) GetAll(c *gin.Context) {
	var tasks []models.Task
	handler.Db.Find(&tasks)
	c.HTML(http.StatusOK, "index.html", gin.H{"tasks": tasks})
}

// 新規作成
func (handler *TaskHandler) Create(c *gin.Context) {
	text, _ := c.GetPostForm("text")
	handler.Db.Create(&models.Task{Text: text})
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (handler *TaskHandler) Edit(c *gin.Context) {

}

func (handler *TaskHandler) Update(c *gin.Context) {

}

// 削除
func (handler *TaskHandler) Delete(c *gin.Context) {
	task := models.Task{}
	id := c.Param("id")
	handler.Db.First(&task, id)
	handler.Db.Delete(&task)
	c.Redirect(http.StatusMovedPermanently, "/")
}
