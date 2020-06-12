package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"strconv"
	"time"
)

type Todo struct {
	ID		  int
	Text      string
	Status    string
	CreatedAt time.Time
}

var Db *gorm.DB

func dbInit() {
	Db, err := gorm.Open("postgres", "user=tomoyaueno dbname=gogin password=gogin sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Todo{})
	defer Db.Close()
}

func dbInsert(text string, status string) {
	Db, err := gorm.Open("postgres", "user=tomoyaueno dbname=gogin password=gogin sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.Create(&Todo{Text: text, Status: status})
	defer Db.Close()
}

func dbUpdate(id int, text string, status string) {
	Db, err := gorm.Open("postgres", "user=tomoyaueno dbname=gogin password=gogin sslmode=disable")
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
	Db, err := gorm.Open("postgres", "user=tomoyaueno dbname=gogin password=gogin sslmode=disable")
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
	Db, err := gorm.Open("postgres", "user=tomoyaueno dbname=gogin password=gogin sslmode=disable")
	if err != nil {
		panic(err)
	}
	var todos []Todo
	Db.Order("created_at desc").Find(&todos)
	Db.Close()
	return todos
}

func dbGetOne(id int) Todo {
	Db, err := gorm.Open("postgres", "user=tomoyaueno dbname=gogin password=gogin sslmode=disable")
	if err != nil {
		panic(err)
	}
	var todo Todo
	Db.First(&todo, id)
	Db.Close()
	return todo
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	//Index
	router.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	//Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)
		ctx.Redirect(302, "/")
	})

	//Show
	router.GET("/show/:id", func(ctx *gin.Context) {
		i := ctx.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "show.html", gin.H{"todo": todo})
	})

	//Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		i := ctx.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbUpdate(id, text, status)
		ctx.Redirect(302, "/")
	})

	//DeleteCheck
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		i := ctx.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	//Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		i := ctx.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		dbDelete(id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
