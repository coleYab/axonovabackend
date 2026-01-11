package usecase

import (
	"axonova/internal/event/dto"
	"axonova/internal/event/entity"
	"axonova/internal/event/repository"
	"axonova/pkg/mailer"
	"axonova/pkg/payment"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v84"
)

type EventUseCase struct {
	gMailer        *mailer.AppMailer
	paymentService payment.PaymentService
	repo           repository.IEventRepository
}

func NewEventUseCase(repo repository.IEventRepository, gMailer *mailer.AppMailer) *EventUseCase {
	return &EventUseCase{repo: repo, gMailer: gMailer}
}

func (au *EventUseCase) HandlePayment(payload []byte, sigHeader string) error {
	const endpointSecret = ""
	event, err := au.paymentService.HandleWebhook(payload, sigHeader, endpointSecret)
	if err != nil {
		return err
	}

	switch event.Type {
	case "checkout.session.completed":
		{
			var session stripe.CheckoutSession
			err := json.Unmarshal(payload, &session)
			if err != nil {
				return err
			}
		}
	case "payment_intent.payment_failed":
		{
			// set it as payment is failed so send the email
		}
	}

	return nil
}

func (au *EventUseCase) GetEventByID(id string) (entity.Event, error) {
	return au.repo.FindByID(id)
}

func (au *EventUseCase) BookEvent(dto dto.BookEventDTO) (payment.CheckoutResult, error) {
	eventID := dto.EventID
	email := dto.Email

	event, err := au.GetEventByID(eventID)
	if err != nil {
		return payment.CheckoutResult{}, err
	}

	paymentResult, err := au.paymentService.CreateCheckoutSession(email, int64(event.Price), event.Title)
	if err != nil {
		return payment.CheckoutResult{}, err
	}

	return paymentResult, nil
}

func (au *EventUseCase) GetAllEvents() ([]entity.Event, error) {
	return au.repo.FindAll()
}

func (au *EventUseCase) DeleteEvent(id string) error {
	return au.repo.Delete(id)
}

func (au *EventUseCase) CreateEvent(createEventDTO dto.CreateEventDTO) (entity.Event, error) {
	event := entity.Event{
		ID:           uuid.NewString(),
		Title:        createEventDTO.Title,
		Picture:      createEventDTO.Picture,
		Description:  createEventDTO.Description,
		Date:         createEventDTO.Date,
		StartTime:    createEventDTO.StartTime,
		DurationMin:  createEventDTO.DurationMin,
		Price:        createEventDTO.Price,
		MaxAttendees: createEventDTO.MaxAttendees,
		IsOnline:     createEventDTO.IsOnline,
		Platform:     createEventDTO.Platform,
		MeetingLink:  createEventDTO.MeetingLink,
		Tags:         createEventDTO.Tags,
	}

	if err := au.repo.Create(event); err != nil {
		return entity.Event{}, err
	}

	return event, nil
}
