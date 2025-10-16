package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ajay-1134/alumni-backend/internal/auth"
	"github.com/ajay-1134/alumni-backend/internal/dto"
	"github.com/ajay-1134/alumni-backend/internal/ports/handler"
	"github.com/ajay-1134/alumni-backend/internal/ports/service"
	"github.com/ajay-1134/alumni-backend/pkg/oauth"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) handler.UserHandler {
	return &userHandler{service: service}
}

func (uh *userHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.service.Register(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

func (uh *userHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}

func (uh *userHandler) GoogleLogin(c *gin.Context) {
	url := oauth.GetGoogleAuthConfig().AuthCodeURL("randomstate")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (uh *userHandler) GoogleCallback(c *gin.Context) {
	code := c.Query("code")

	token, err := oauth.GetGoogleAuthConfig().Exchange(context.Background(), code)
	if err != nil {
		log.Println("code exchange failed: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to exhange code", "details": err.Error()})
		return
	}

	client := oauth.GetGoogleAuthConfig().Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get user info"})
		return
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	var googleUser dto.GoogleUser

	if err := json.Unmarshal(data, &googleUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse user info"})
		return
	}

	user, err := uh.service.LoginWithGoogle(&googleUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := auth.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/auth/callback?token="+tokenString)
}

func (uh *userHandler) GetAllUsers(c *gin.Context) {
	users, err := uh.service.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uh *userHandler) GetMyProfile(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		log.Printf("error occured in getting the id in the handler")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("id while getting user profile: %v\n", id)
	user, err := uh.service.GetUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *userHandler) UpdateUserDetails(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var req dto.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("error occured in binding the gin context!")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uh.service.UpdateDetails(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user details updated successfully!"})
}

func (uh *userHandler) DeleteProfile(c *gin.Context) {
	id, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = uh.service.DeleteProfile(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user profile deleted successfully"})
}

func(uh *userHandler) GetTotalUserCount(c *gin.Context) {
	totalUsers, err := uh.service.UserCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK,&totalUsers)
}

func(uh *userHandler) GetTotalVerifiedUsersCount(c *gin.Context) {
	totalVerifiedUsers, err := uh.service.VerifiedUsersCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		return 
	}

	c.JSON(http.StatusOK,&totalVerifiedUsers)
}
