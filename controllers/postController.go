package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/n-had/golang-crud-boilerplate/initializers"
	"github.com/n-had/golang-crud-boilerplate/models"
)

func CreatePost(ctx *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	ctx.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPosts(ctx *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostById(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post

	initializers.DB.Find(&post, id)

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var body struct {
		Title string
		Body  string
	}
	ctx.Bind(&body)
	var post models.Post

	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	ctx.Status(200)
}
