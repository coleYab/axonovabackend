package entity

type Assessment struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company,omitempty"`

	Answers    map[string]int `json:"answers"`
	TotalScore int            `json:"totalScore"`

	RecommendationTitle string `json:"recommendationTitle"`

	AnsweredCount  int `json:"answeredCount"`
	TotalQuestions int `json:"totalQuestions"`
}
