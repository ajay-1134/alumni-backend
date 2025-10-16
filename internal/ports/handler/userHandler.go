package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GoogleLogin(c *gin.Context)
	GoogleCallback(c *gin.Context)
	GetAllUsers(c *gin.Context)
	GetMyProfile(c *gin.Context)
	UpdateUserDetails(c *gin.Context)
	DeleteProfile(c *gin.Context)
	GetTotalUserCount(c *gin.Context)
	GetTotalVerifiedUsersCount(c *gin.Context)
}
