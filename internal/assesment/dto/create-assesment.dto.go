package dto

type CreateAssessmentDTO struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Company string `json:"company,omitempty"`

	Answers    map[string]int `json:"answers" binding:"required"`
	TotalScore int            `json:"totalScore" binding:"required,min=0,max=100"`

	RecommendationTitle string `json:"recommendationTitle" binding:"required"`

	AnsweredCount  int `json:"answeredCount" binding:"required,min=0"`
	TotalQuestions int `json:"totalQuestions" binding:"required,min=1"`
}
