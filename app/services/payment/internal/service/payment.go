package service

import (
	"context"
	"log"
	"monolith-divided-to-microservices/app/services/payment/internal/config"
	"monolith-divided-to-microservices/app/services/payment/internal/repository/postgres"
	"monolith-divided-to-microservices/app/services/payment/internal/schema"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

type Services struct {
	PaymentService *PaymentService
}

func (s *PaymentService) CreatePayment(ctx context.Context, req *schema.SavePaymentRequest) (*stripe.CheckoutSession, error) {
	stripe.Key = s.cfg.StripeSecret
	amountInCents := int64(req.Amount * 100)
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(string(stripe.CurrencyUSD)),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Order #" + req.OrderID),
					},
					UnitAmount: stripe.Int64(amountInCents),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(s.cfg.Domain + "/success?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(s.cfg.Domain + "/cancel"),
	}

	params.AddMetadata("user_id", req.UserID)
	params.AddMetadata("order_id", req.OrderID)

	sess, err := session.New(params)
	if err != nil {
		log.Printf("Error creating session: %v", err)
		return nil, err
	}

	return sess, nil
}

func (s *PaymentService) GetPaymentStatus(sessionID string) (*stripe.CheckoutSession, error) {
	stripe.Key = s.cfg.StripeSecret
	session, err := session.Get(sessionID, nil)
	if err != nil {
		log.Printf("error retrieving session: %v", err)
		return nil, err
	}

	return session, nil
}

func NewPaymentService(repo *postgres.PaymentRepository, cfg *config.Config) *PaymentService {
	return &PaymentService{
		repo: repo,
		cfg:  cfg,
	}
}
