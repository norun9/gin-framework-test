package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model

}

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	data := "Hello Go/Gin"
	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.Run()
}
