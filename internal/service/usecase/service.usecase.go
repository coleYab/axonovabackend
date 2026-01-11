package usecase

import (
	"axonova/internal/mailist"
	"axonova/internal/service/dto"
	"axonova/internal/service/entity"
	"axonova/internal/service/repository"
	"axonova/pkg/mailer"

	"github.com/google/uuid"
)

type ServiceUseCase struct {
	gMailer *mailer.AppMailer
	repo    repository.IServiceRepository
}

func NewServiceUseCase(repo repository.IServiceRepository, gMailer *mailer.AppMailer) *ServiceUseCase {
	return &ServiceUseCase{repo: repo, gMailer: gMailer}
}

func (su *ServiceUseCase) CreateServiceRequest(createServiceRequest dto.ServiceRequestDTO) (entity.Service, error) {
	service := entity.Service{
		ID:               uuid.NewString(),
		Name:             createServiceRequest.Name,
		Email:            createServiceRequest.Email,
		Message:          createServiceRequest.Message,
		Service:          createServiceRequest.Service,
		Phone:            createServiceRequest.Phone,
		RequestedModules: createServiceRequest.GetModulesSlice(),
		PreferredDate:    createServiceRequest.PreferredDate,
	}

	if err := su.repo.CreateService(service); err != nil {
		return entity.Service{}, err
	}

	acknowledgment, err := mailist.GenerateSenderAcknowledgment(service.Name, true)
	if err != nil {
		return entity.Service{}, err
	}

	if err := su.gMailer.SendGmail(service.Email, "New axonova consulting service", acknowledgment); err != nil {
		return entity.Service{}, err
	}

	acknowledgment, err = mailist.GenerateServiceRequestEmail(mailist.ServiceRequestData{
		Name:             service.Name,
		Email:            service.Email,
		Message:          service.Message,
		Phone:            service.Phone,
		ServiceType:      service.Service,
		PreferredDate:    service.PreferredDate,
		RequestedModules: service.RequestedModules,
	})
	if err != nil {
		return entity.Service{}, err
	}
	if err := su.gMailer.SendGmail(service.Email, "New service request | axonova consulting", acknowledgment); err != nil {
		return entity.Service{}, err
	}

	return service, nil
}

func (su *ServiceUseCase) CreateContactRequest(createContactRequest dto.ContactRequestDTO) (entity.Contact, error) {
	contactRequest := entity.Contact{
		ID:      uuid.NewString(),
		Name:    createContactRequest.Name,
		Email:   createContactRequest.Email,
		Message: createContactRequest.Message,
	}

	if err := su.repo.CreateContact(contactRequest); err != nil {
		return entity.Contact{}, err
	}

	acknowledgment, err := mailist.GenerateSenderAcknowledgment(contactRequest.Name, false)
	if err != nil {
		return entity.Contact{}, err
	}

	if err := su.gMailer.SendGmail(contactRequest.Email, "Axonova consulting service contactus", acknowledgment); err != nil {
		return entity.Contact{}, err
	}

	acknowledgment, err = mailist.GenerateContactFormEmail(contactRequest.Name, contactRequest.Email, contactRequest.Message)
	if err != nil {
		return entity.Contact{}, err
	}
	if err := su.gMailer.SendGmail(contactRequest.Email, "Contact US form submission | axonova consulting", acknowledgment); err != nil {
		return entity.Contact{}, err
	}

	return contactRequest, nil
}
