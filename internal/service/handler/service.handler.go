package handler

import (
	"axonova/internal/service/dto"
	"axonova/internal/service/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	auc *usecase.ServiceUseCase
}

func NewServiceHandler(auc *usecase.ServiceUseCase) *ServiceHandler {
	return &ServiceHandler{
		auc: auc,
	}
}

func (ah *ServiceHandler) CreateService(ctx *gin.Context) {
	var serviceDTO dto.ServiceRequestDTO
	if err := ctx.ShouldBindJSON(&serviceDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service, err := ah.auc.CreateServiceRequest(serviceDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"service": service})
}

func (ah *ServiceHandler) CreateContact(ctx *gin.Context) {
	var contactDTO dto.ContactRequestDTO
	if err := ctx.ShouldBindJSON(&contactDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact, err := ah.auc.CreateContactRequest(contactDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"contact": contact})
}

func (ah *ServiceHandler) RegisterRoutes(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/contact", ah.CreateContact)
	routeGroup.POST("/service", ah.CreateService)
}
