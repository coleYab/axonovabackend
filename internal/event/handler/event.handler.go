package handler

import (
	"axonova/internal/event/dto"
	"axonova/internal/event/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	auc *usecase.EventUseCase
}

func NewEventHandler(auc *usecase.EventUseCase) *EventHandler {
	return &EventHandler{
		auc: auc,
	}
}

func (ah *EventHandler) CreateEvent(ctx *gin.Context) {
	var createEventDTO dto.CreateEventDTO
	if err := ctx.ShouldBindJSON(&createEventDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := ah.auc.CreateEvent(createEventDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"event": event})
}

func (ah *EventHandler) DeleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := ah.auc.DeleteEvent(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (ah *EventHandler) FindEventByID(ctx *gin.Context) {
	id := ctx.Param("id")
	event, err := ah.auc.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"event": event})
}

func (ah *EventHandler) BookEvent(ctx *gin.Context) {
	var bookEventDTO dto.BookEventDTO
	if err := ctx.ShouldBindJSON(&bookEventDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookingData, err := ah.auc.BookEvent(bookEventDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"result": bookingData})
}

func (h *EventHandler) StripeWebhookHandler(ctx *gin.Context) {
	payload, err := ctx.GetRawData()
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	sigHeader := ctx.GetHeader("Stripe-Signature")
	if err := h.auc.HandlePayment(payload, sigHeader); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusOK)
}

func (ah *EventHandler) FindAllEvents(ctx *gin.Context) {
	events, err := ah.auc.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"events": events})
}

func (ah *EventHandler) RegisterRoutes(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/:id", ah.FindEventByID)
	routeGroup.POST("/", ah.CreateEvent)
	routeGroup.GET("/", ah.FindAllEvents)
	routeGroup.DELETE("/:id", ah.DeleteEvent)
}
