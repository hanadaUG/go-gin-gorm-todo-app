package main

import (
	"github.com/hanadaUG/go-gin-gorm-todo-app/db"
	"github.com/hanadaUG/go-gin-gorm-todo-app/router"
)

func main() {
	// DBのOpen & Close処理の遅延実行
	db.Initialize()
	defer db.Close()

	// ルーティング設定
	router.Router()
}
