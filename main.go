package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

func dbInit() {
	db, err := sql.Open("postgres", "username=tomoyaueno dbname=gogin password=gogin sslmode=disable")
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
