package schema

type SavePaymentRequest struct {
	UserID  string  `json:"user_id" binding:"required"`
	OrderID string  `json:"order_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
	Method  string  `json:"method" binding:"required"` // card, cash, etc
}
