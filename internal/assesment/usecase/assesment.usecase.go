package usecase

import (
	"axonova/internal/assesment/dto"
	"axonova/internal/assesment/entity"
	"axonova/internal/assesment/repository"
	"axonova/pkg/mailer"
	"fmt"

	"github.com/google/uuid"
)

type AssessmentUseCase struct {
	gMailer *mailer.AppMailer
	repo    repository.IAssessmentRepository
}

func NewAssessmentUseCase(repo repository.IAssessmentRepository, gMailer *mailer.AppMailer) *AssessmentUseCase {
	return &AssessmentUseCase{repo: repo, gMailer: gMailer}
}

func (au *AssessmentUseCase) GetAssessmentByID(id string) (entity.Assessment, error) {
	return au.repo.FindByID(id)
}

func (au *AssessmentUseCase) GetAllAssessment() ([]entity.Assessment, error) {
	return au.repo.FindAll()
}

func (au *AssessmentUseCase) DeleteAssessment(id string) error {
	return au.repo.Delete(id)
}

func (au *AssessmentUseCase) CreateAssessment(assessmentDto dto.CreateAssessmentDTO) (entity.Assessment, error) {
	assessment := entity.Assessment{
		ID:                  uuid.NewString(),
		Name:                assessmentDto.Name,
		Email:               assessmentDto.Email,
		Company:             assessmentDto.Company,
		Answers:             assessmentDto.Answers,
		TotalScore:          assessmentDto.TotalScore,
		RecommendationTitle: assessmentDto.RecommendationTitle,
		AnsweredCount:       assessmentDto.AnsweredCount,
		TotalQuestions:      assessmentDto.TotalQuestions,
	}
	if err := au.repo.Create(assessment); err != nil {
		return entity.Assessment{}, err
	}

	if err := au.gMailer.SendGmail(assessment.Email, "You have the user", "<p>Who am I</p>"); err != nil {
		return entity.Assessment{}, fmt.Errorf("failed to send gmail email: %v", err)
	}

	return assessment, nil
}
