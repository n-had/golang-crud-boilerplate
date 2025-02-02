package main

import (
	"github.com/gin-gonic/gin"
	"github.com/n-had/golang-crud-boilerplate/controllers"
	"github.com/n-had/golang-crud-boilerplate/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetPostById)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run()
}
