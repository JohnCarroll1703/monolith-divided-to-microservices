package v1

import (
	"monolith-divided-to-microservices/app/services/user/internal/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) login(c *gin.Context) {
	var req schema.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.Services.UserService.GetByEmailAndPassword(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := h.JWT.Generate(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, schema.LoginResponse{Token: token})
}
