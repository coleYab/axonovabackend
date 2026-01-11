package dto

import "strings"

type ServiceRequestDTO struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`

	Phone            string `json:"phone,omitempty"`
	Service          string `json:"service,omitempty"`
	PreferredDate    string `json:"preferred_date,omitempty"`
	RequestedModules string `json:"requested_modules,omitempty"`
}

func (d *ServiceRequestDTO) GetModulesSlice() []string {
	if d.RequestedModules == "" {
		return nil
	}
	parts := strings.Split(d.RequestedModules, ",")
	var cleaned []string
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			cleaned = append(cleaned, trimmed)
		}
	}
	return cleaned
}
