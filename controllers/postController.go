package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/notoriousCode97/go-crud/initializers"
	"github.com/notoriousCode97/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with the posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Update the post
	post.Title = body.Title
	post.Body = body.Body

	initializers.DB.Save(&post)

	// Respond with the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Delete the post
	initializers.DB.Delete(&post)

	// Respond with the post
	c.Status(200)
}
