package router

import (
	"github.com/ajay-1134/alumni-backend/internal/handler"
	"github.com/ajay-1134/alumni-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, uh *handler.UserHandler) {
	r.POST("users/register", uh.Register)
	r.POST("users/login", uh.Login)

	r.GET("auth/google/login", uh.GoogleLogin)
	r.GET("auth/google/callback", uh.GoogleCallback)

	userControl := r.Group("users")
	userControl.Use(middleware.AuthMiddleware())
	{
		userControl.PATCH("update/me", uh.UpdateUserDetails)
		userControl.DELETE("delete/me", uh.DeleteProfile)
		userControl.GET("/me", uh.GetMyProfile)
	}

	adminControl := r.Group("users")
	adminControl.Use(middleware.AdminOnly())
	{
		adminControl.GET("all", uh.GetAllUsers)
		adminControl.PATCH("update/:id", uh.UpdateUserDetails)
		adminControl.DELETE("delete/:id", uh.DeleteProfile)
		adminControl.GET("/:id", uh.GetMyProfile)
	}

}
