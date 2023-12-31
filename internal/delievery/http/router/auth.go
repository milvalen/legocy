package v1

import (
	"github.com/gin-gonic/gin"
	_ "legocy-go/docs"
	"legocy-go/internal/delievery/http/handlers/users/auth"
	m "legocy-go/internal/delievery/http/middleware"
	s "legocy-go/internal/domain/users/service"
)

func (r V1router) addAuth(rg *gin.RouterGroup, service s.UserUseCase) {
	handler := auth.NewTokenHandler(service)

	authRouter := rg.Group("/users")
	{
		authRouter.POST("/token", handler.GenerateToken)
		authRouter.POST("/register", handler.UserRegister)
	}

	authPrivate := rg.Group("/admin/users").Use(m.AdminUserOnly())
	{
		authPrivate.POST("/", handler.AdminRegister)
	}
}
