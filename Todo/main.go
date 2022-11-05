package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/Todo")
	{
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}
	router.Run()
}
var db *gorm.DB

func init () {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:12345@/Pelumi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&todoModel{})
}

type todoModel struct {
	gorm.Model
	Title	string `json:"title"`
	Completed	int `json:"completed"`
}

type transformedTodo struct{
	ID	unit `json:"id"`
	Title	string `json:"title"`
	Completed 	bool `json:"completed"`
}

func createTodo(c *gin.Context){
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := todoModel{Title: c.PostForm("title"), Completed: completed}
	db.Save(&todo)
	c.JSON(http.StatusCreated,gin.H{"status": http.StatusCreated, "message":"Todo item created successfully!", "resourceId":todo.ID})
}
func fetchAllTodo