package controller

import (
	"fmt"
	"log"

	"github.com/LuanTruongPTIT/tutor-be/modules/auth/dto"
	"github.com/LuanTruongPTIT/tutor-be/modules/auth/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.IAuthService
}

func NewUserHandler(service services.IAuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		log.Fatal("Failed to get body", err)
		//  response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
	}

	fmt.Print("fail r", &req)
}
