package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hanadaUG/go-gin-gorm-todo-app/controllers"
	"github.com/hanadaUG/go-gin-gorm-todo-app/db"
)

func Router() {
	// gin内で定義されているEngine構造体インスタンスを取得
	// Router、HTML Rendererの機能を内包している
	router := gin.Default()

	// globパターンに一致するHTMLファイルをロードしHTML Rendererに関連付ける
	// 今回のケースではtemplatesディレクトリ配下のhtmlファイルを関連付けている
	router.LoadHTMLGlob("templates/*.html")

	// TaskHandler構造体に紐付けたCRUDメソッドを呼び出す
	handler := controllers.TaskHandler{
		db.Get(),
	}

	// 各パスにGET/POSTメソッドでリクエストされた時のレスポンスを登録する
	// 第一引数にパス、第二引数にHandlerを登録する
	router.GET("/", handler.GetAll)            // 一覧表示
	router.POST("/", handler.Create)           // 新規作成
	router.POST("/delete/:id", handler.Delete) // 削除

	// Routerをhttp.Serverに接続し、HTTPリクエストのリスニングとサービスを開始する
	router.Run()
}
