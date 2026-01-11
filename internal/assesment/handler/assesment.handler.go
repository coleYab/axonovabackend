package handler

import (
	"axonova/internal/assesment/dto"
	"axonova/internal/assesment/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssessmentHandler struct {
	auc *usecase.AssessmentUseCase
}

func NewAssessmentHandler(auc *usecase.AssessmentUseCase) *AssessmentHandler {
	return &AssessmentHandler{
		auc: auc,
	}
}

func (ah *AssessmentHandler) CreateAssessment(ctx *gin.Context) {
	var createAssessmentDto dto.CreateAssessmentDTO
	if err := ctx.ShouldBindJSON(&createAssessmentDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assessment, err := ah.auc.CreateAssessment(createAssessmentDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"assessment": assessment})
}

//func (ah *AssessmentHandler) UpdateAssessment(ctx *gin.Context) {
//	id := ctx.Param("id")
//	//assessment, err := ah.auc.GetAssessmentByID(id)
//}

func (ah *AssessmentHandler) DeleteAssessment(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := ah.auc.DeleteAssessment(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})

}
func (ah *AssessmentHandler) FindAssessmentByID(ctx *gin.Context) {
	id := ctx.Param("id")
	assessment, err := ah.auc.GetAssessmentByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"assessment": assessment})
}

func (ah *AssessmentHandler) FindAllAssessments(ctx *gin.Context) {
	assessments, err := ah.auc.GetAllAssessment()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"assessments": assessments})
}

func (ah *AssessmentHandler) RegisterRoutes(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/:id", ah.FindAssessmentByID)
	routeGroup.POST("/", ah.CreateAssessment)
	routeGroup.GET("/", ah.FindAllAssessments)
	routeGroup.DELETE("/:id", ah.DeleteAssessment)
}
