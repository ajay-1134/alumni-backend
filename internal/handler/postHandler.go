package handler

import (
	"log"
	"net/http"

	"github.com/ajay-1134/alumni-backend/internal/dto"
	"github.com/ajay-1134/alumni-backend/internal/ports/handler"
	"github.com/ajay-1134/alumni-backend/internal/ports/service"
	"github.com/gin-gonic/gin"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) handler.PostHandler {
	return &postHandler{service: service}
}

func (ph *postHandler) CreatePost(c *gin.Context) {
	userID,err := getUserID(c)

	if err != nil {
		log.Printf("error occured in finding the user id")
		c.JSON(http.StatusUnauthorized,gin.H{"error" : err.Error()})
		return
	}

	var postDto *dto.PostDto 

	if err := c.ShouldBindJSON(&postDto); err != nil {
		log.Printf("error occured in binding request into post")
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
		return
	}

	postDto.UserID = userID

	err = ph.service.CreatePost(postDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		return 
	}

	c.JSON(http.StatusCreated,gin.H{"message" : "post created successfully"})
}

func (ph *postHandler) GetAllPosts(c *gin.Context) {
	posts, err := ph.service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		return 
	}

	c.JSON(http.StatusOK,&posts)
}

func (ph *postHandler) GetAllPostsWithUserId(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		log.Printf("error occured in getting the id in the handler")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	posts,err := ph.service.GetAllPostsWithUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK,&posts)
}

func (ph *postHandler) UpdatePost(c *gin.Context) {
	id := c.Param("postID")
	postID, _ := stringToUint(id)
	var postDto *dto.PostDto 

	if err := c.ShouldBindJSON(&postDto); err != nil {
		log.Printf("error occured in binding request into post")
		c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
		return
	}

	err := ph.service.UpdatePost(postID,postDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message" : "post updated successfully"})
}

func (ph *postHandler) DeletePost(c *gin.Context) {
	id := c.Param("postID")
	postID, _ := stringToUint(id)

	err := ph.service.DeletePost(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message" : "post deleted successfully"})
}

