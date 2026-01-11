package handler

import (
	"axonova/internal/event/dto"
	"axonova/internal/event/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	auc *usecase.EventUseCase
}

func NewServiceHandler(auc *usecase.EventUseCase) *ServiceHandler {
	return &ServiceHandler{
		auc: auc,
	}
}

func (ah *ServiceHandler) CreateEvent(ctx *gin.Context) {
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

func (ah *ServiceHandler) DeleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := ah.auc.DeleteEvent(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (ah *ServiceHandler) FindEventByID(ctx *gin.Context) {
	id := ctx.Param("id")
	event, err := ah.auc.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"event": event})
}

func (ah *ServiceHandler) BookEvent(ctx *gin.Context) {
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

func (ah *ServiceHandler) RegisterRoutes(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/:id", ah.FindEventByID)
	routeGroup.POST("/", ah.CreateEvent)
	routeGroup.DELETE("/:id", ah.DeleteEvent)
}
