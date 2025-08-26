package v1

import (
	"monolith-divided-to-microservices/app/services/item/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Logger   *logrus.Logger
	Services *service.Services
}

func NewHandler(logger *logrus.Logger,
	services *service.Services) *Handler {
	return &Handler{
		Logger:   logger,
		Services: services,
	}
}

func (h *Handler) Health(c *gin.Context) {
	h.Logger.Info("Health check called")
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
