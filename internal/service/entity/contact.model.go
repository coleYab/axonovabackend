package entity

type Contact struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Service struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Email            string   `json:"email"`
	Message          string   `json:"message"`
	Phone            string   `json:"phone,omitempty"`
	Service          string   `json:"service"`
	PreferredDate    string   `json:"preferred_date"`
	RequestedModules []string `json:"requested_modules"`
}
