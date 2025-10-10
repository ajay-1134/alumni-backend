package router

import (
	"github.com/ajay-1134/alumni-backend/internal/middleware"
	"github.com/ajay-1134/alumni-backend/internal/ports/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutesforUser(r *gin.Engine, uh handler.UserHandler) {
	r.POST("auth/register", uh.Register)
	r.POST("auth/login", uh.Login)

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

func SetupRoutesforPost(r *gin.Engine, ph handler.PostHandler) {
	userControl := r.Group("feed")
	userControl.Use(middleware.AuthMiddleware())
	{
		userControl.GET("",ph.GetAllPosts)
		userControl.POST("post",ph.CreatePost)
		userControl.PATCH("/post/edit/:postID",ph.UpdatePost)
		userControl.GET("user/posts",ph.GetAllPostsWithUserId)
		userControl.DELETE("post/delete/:postID",ph.DeletePost)
	}
}