package router

import (
	"github.com/LuanTruongPTIT/tutor-be/modules/Auth/port/controller"
	"github.com/LuanTruongPTIT/tutor-be/modules/Auth/repository"
	"github.com/LuanTruongPTIT/tutor-be/modules/Auth/services"
	"github.com/LuanTruongPTIT/tutor-be/pkg/dbs"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, sqlDB dbs.IDatabase) {
	authRepo := repository.NewAuthRepository(sqlDB)
	authSvc := services.NewAuthService(authRepo)
	authHandler := controller.NewUserHandler(authSvc)
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", authHandler.Register)
	}
}
