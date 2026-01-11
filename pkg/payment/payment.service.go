package payment

import (
	"fmt"

	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/checkout/session"
	"github.com/stripe/stripe-go/v84/webhook"
)

type CheckoutResult struct {
	SessionID  string `json:"sessionId"`
	SessionURL string `json:"sessionUrl"`
}

type PaymentService interface {
	CreateCheckoutSession(email string, amount int64, eventTitle string) (CheckoutResult, error)
	HandleWebhook(payload []byte, sigHeader string, webhookSecret string) (*stripe.Event, error)
}

type stripeService struct {
	successURL string
	cancelURL  string
	apiKey     string
}

func NewStripeService(apiKey, successURL, cancelURL string) PaymentService {
	stripe.Key = apiKey
	return &stripeService{apiKey: apiKey, successURL: successURL, cancelURL: cancelURL}
}

func (s *stripeService) HandleWebhook(payload []byte, sigHeader string, webhookSecret string) (*stripe.Event, error) {
	// Verify the event signature
	event, err := webhook.ConstructEvent(payload, sigHeader, webhookSecret)
	if err != nil {
		return nil, fmt.Errorf("invalid signature: %v", err)
	}

	return &event, nil
}

func (s *stripeService) CreateCheckoutSession(email string, amount int64, eventTitle string) (CheckoutResult, error) {
	params := &stripe.CheckoutSessionParams{
		CustomerEmail: stripe.String(email),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("gbp"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(eventTitle),
					},
					UnitAmount: stripe.Int64(amount),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(s.successURL),
		CancelURL:  stripe.String(s.cancelURL),
	}

	sess, err := session.New(params)
	if err != nil {
		return CheckoutResult{}, err
	}

	return CheckoutResult{
		SessionID:  sess.ID,
		SessionURL: sess.URL,
	}, nil
}
