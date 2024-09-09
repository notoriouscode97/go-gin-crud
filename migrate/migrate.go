package main

import (
	"github.com/notoriousCode97/go-crud/initializers"
	"github.com/notoriousCode97/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
