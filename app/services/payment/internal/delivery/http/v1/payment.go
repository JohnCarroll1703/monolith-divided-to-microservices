package v1

import (
	"monolith-divided-to-microservices/app/services/payment/internal/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) savePayment(c *gin.Context) {
	var req schema.SavePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.Services.PaymentService.CreatePayment(c.Request.Context(), &req)
	if err != nil {
		h.Logger.Errorf("error creating payment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create payment"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"payment": payment})

}

func (h *Handler) getPaymentStatus(c *gin.Context) {
	sessionId := c.Query("session_id")
	if sessionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id is required"})
		return
	}

	payment, err := h.Services.PaymentService.GetPaymentStatus(sessionId)
	if err != nil {
		h.Logger.Errorf("error retrieving payment status: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve payment status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}
