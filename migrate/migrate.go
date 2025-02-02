package main

import (
	"github.com/n-had/golang-crud-boilerplate/initializers"
	"github.com/n-had/golang-crud-boilerplate/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
