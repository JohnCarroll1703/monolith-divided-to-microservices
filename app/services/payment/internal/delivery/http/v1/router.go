package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Init() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.POST("/payment", h.savePayment)
	v1.GET("/payment/status/:id", h.getPaymentStatus)

	return r
}
