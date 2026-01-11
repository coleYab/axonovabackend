package dto

type ContactRequestDTO struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
