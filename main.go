package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Todo struct {
	Text   string
	Status string
}

var Db *gorm.DB

func dbInit() {
	Db, err := gorm.Open("postgres", "username=tomoyaueno dbname=gogin password=gogin sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Todo{})
	defer Db.Close()
}

func dbInsert(text string, status string) {
	Db, err := gorm.Open("postgres", "username=tomoyaueno dbname=gogin password=gogin sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.Create(&Todo{Text: text, Status: status})
	defer Db.Close()
}

func dbUpdate(id int, text string, status string) {
	Db, err := gorm.Open("postgres", "username=tomoyaueno dbname=gogin password=gogin sslmode=disable")
	if err != nil {
		panic(err)
	}
	var todo Todo
	Db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	Db.Save(&todo)
	Db.Close()
}

func dbDelete(id int) {
	Db, err := gorm.Open("postgres", "username=tomoyaueno dbname=gogin password=gogin sslmode=disable")
	if err != nil {
		panic(err)
	}
	var todo Todo
	//引数のidを参照して最初のレコードを取得
	Db.First(&todo, id)
	Db.Delete(&todo)
	Db.Close()
}

func dbGetAll() []Todo {
	Db, err := gorm.Open("postgres", "username=tomoyaueno dbname=gogin password=gogin sslmode=disable")
	if err != nil {
		panic(err)
	}
	var todos []Todo
	Db.Order("created_at desc").Find(&todos)
	Db.Close()
	return todos
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	data := "Hello Go/Gin"
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.Run()
}
