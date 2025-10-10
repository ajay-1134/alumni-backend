package handler

import "github.com/gin-gonic/gin"

type PostHandler interface {
	CreatePost(c *gin.Context)
	GetAllPosts(c *gin.Context)
	GetAllPostsWithUserId(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}