package app

import (
	"github.com/ajay-1134/alumni-backend/internal/handler"
	"github.com/ajay-1134/alumni-backend/internal/repository"
	"github.com/ajay-1134/alumni-backend/internal/router"
	"github.com/ajay-1134/alumni-backend/internal/service"
	"github.com/ajay-1134/alumni-backend/pkg/db"
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

	UserService := service.NewUserService(userRepository)

	UserHandler := handler.NewUserHandler(UserService)

	router.SetupRoutes(b.router, UserHandler)

	return b.router
}
