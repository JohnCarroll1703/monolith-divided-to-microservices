package v1

import (
	"monolith-divided-to-microservices/app/services/user/internal/schema"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.Services.UserService.GetAllUsers(c.Request.Context())
	if err != nil {
		h.Logger.Error("error fetching users: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(200, users)
}

func (h *Handler) createUser(c *gin.Context) {
	var req schema.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	if err := h.Services.UserService.CreateUser(c.Request.Context(), req); err != nil {
		h.Logger.Errorf("error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": req})
}

func (h *Handler) getUser(c *gin.Context) {
	var filters schema.UserFilters
	filters.ID = c.Query("id")
	filters.Username = c.Query("username")
	filters.Email = c.Query("email")

	if cb := c.Query("created_before"); cb != "" {
		t, _ := time.Parse("2006-01-02", cb)
		filters.CreatedBefore = t
	}
	if ca := c.Query("created_after"); ca != "" {
		t, _ := time.Parse("2006-01-02", ca)
		filters.CreatedAfter = t
	}
	users, err := h.Services.UserService.GetUser(c.Request.Context(), filters)
	if err != nil {
		h.Logger.Errorf("error fetching users: %v", err)
		c.JSON(500, gin.H{"error": "internal error"})
	}
	c.JSON(200, users)
}
