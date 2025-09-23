package app

import (
	"github.com/ajay-1134/alumni-backend/internal/handler"
	"github.com/ajay-1134/alumni-backend/internal/repository"
	"github.com/ajay-1134/alumni-backend/internal/router"
	"github.com/ajay-1134/alumni-backend/internal/service"
	"github.com/ajay-1134/alumni-backend/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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

	// apply CORS before routes
	b.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// now set up routes
	router.SetupRoutes(b.router, userHandler)

	return b.router
}
