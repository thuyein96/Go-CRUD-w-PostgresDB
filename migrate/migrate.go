package main

import (
	"github.com/thuyein/go-crud/initializers"
	"github.com/thuyein/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}