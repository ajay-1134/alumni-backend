package app

import (
	"github.com/ajay-1134/alumni-backend/internal/handler"
	"github.com/ajay-1134/alumni-backend/internal/repository"
	"github.com/ajay-1134/alumni-backend/internal/router"
	"github.com/ajay-1134/alumni-backend/internal/service"
	"github.com/ajay-1134/alumni-backend/pkg/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Builder struct {
	router *gin.Engine
}

func NewBuilder() *Builder {
	return &Builder{router: gin.Default()}
}

func (b *Builder) Build(dsn string) *gin.Engine {
	database := db.ConnnectDB(dsn)

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	b.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 👈 allow cookies
	}))

	// now set up routes
	router.SetupRoutes(b.router, userHandler)

	return b.router
}
