package v1

import (
	"monolith-divided-to-microservices/app/services/user/internal/auth"
	"monolith-divided-to-microservices/app/services/user/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Logger   *logrus.Logger
	Services *service.Services
	JWT      *auth.JWTManager
}

func NewHandler(logger *logrus.Logger,
	services *service.Services,
	jwtManager *auth.JWTManager) *Handler {
	return &Handler{
		Logger:   logger,
		Services: services,
		JWT:      jwtManager,
	}
}

func (h *Handler) Health(c *gin.Context) {
	h.Logger.Info("Health check called")
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
