package server

import (
	assessmentHandler "axonova/internal/assesment/handler"
	assessmentRepository "axonova/internal/assesment/repository"
	assessmentUseCase "axonova/internal/assesment/usecase"
	eventHandler "axonova/internal/event/handler"
	eventRepository "axonova/internal/event/repository"
	eventUseCase "axonova/internal/event/usecase"
	"axonova/pkg/database"
	"axonova/pkg/mailer"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppServer struct {
	engine *gin.Engine
}

func NewAppServer() *AppServer {
	return &AppServer{
		engine: gin.Default(),
	}
}

func (server *AppServer) RegisterRoutes(database *database.MongoDB, gMailer *mailer.AppMailer) {
	apiSubgroup := server.engine.Group("/api")

	// assessment related things
	assessmentCollection := database.GetCollection("assessment")
	assessmentSubGroup := apiSubgroup.Group("/assessment")
	aRepository := assessmentRepository.NewMongoAssessmentRepository(assessmentCollection)
	aUseCase := assessmentUseCase.NewAssessmentUseCase(aRepository, gMailer)
	aHandler := assessmentHandler.NewAssessmentHandler(aUseCase)
	aHandler.RegisterRoutes(assessmentSubGroup)

	// event related things
	eventCollection := database.GetCollection("event")
	eventSubGroup := apiSubgroup.Group("/event")
	eveRepository := eventRepository.NewMongoEventRepository(eventCollection)
	eveUseCase := eventUseCase.NewEventUseCase(eveRepository, gMailer)
	eveHandler := eventHandler.NewEventHandler(eveUseCase)
	eveHandler.RegisterRoutes(eventSubGroup)

	server.engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

func (server *AppServer) Run(addr string) {
	fmt.Println("Starting server on " + addr)
	if err := server.engine.Run(addr); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
