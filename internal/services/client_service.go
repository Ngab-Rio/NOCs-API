package services

import (
	"context"
	"errors"

	"github.com/Ngab-Rio/NOCs-API/internal/dto"
	appErr "github.com/Ngab-Rio/NOCs-API/internal/errors"
	"github.com/Ngab-Rio/NOCs-API/internal/models"
	"github.com/Ngab-Rio/NOCs-API/internal/repository"
	"gorm.io/gorm"
)

type ClientService interface {
	CreateClient(ctx context.Context, req dto.CreateClientRequest) (*dto.ClientResponse, error)
	UpdateClient(ctx context.Context, ClientID int, req dto.UpdateClientRequest) (*dto.ClientResponse, error)
	DeleteClient(ctx context.Context, ClientID int) error
	GetClientByID(ctx context.Context, ClientID int) (*dto.ClientResponse, error)
	GetClients(ctx context.Context, req dto.GetClientsRequest) (*dto.GetClientsResponse, error)
}

type clientService struct {
	clientRepo repository.ClientRepository
}

func NewClientService(clientRepo repository.ClientRepository) ClientService {
	return &clientService{clientRepo: clientRepo}
}

func (s *clientService) CreateClient(ctx context.Context, req dto.CreateClientRequest) (*dto.ClientResponse, error) {
	existingClient, err := s.clientRepo.FindByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingClient != nil {
		return nil, appErr.ErrInvalidRequest
	}

	existingPhone, err := s.clientRepo.FindByPhone(ctx, req.Phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingPhone != nil {
		return nil, appErr.ErrInvalidRequest
	}

	client := &models.Client{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
		Status:    "active",
	}

	if err := s.clientRepo.Create(ctx, client); err != nil {
		return nil, err
	}

	return &dto.ClientResponse{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		Address:   client.Address,
		Longitude: client.Longitude,
		Latitude:  client.Latitude,
		Status:    client.Status,
	}, nil
}

func (s *clientService) UpdateClient(ctx context.Context, clientID int, req dto.UpdateClientRequest) (*dto.ClientResponse, error) {
	client, err := s.clientRepo.FindByID(ctx, clientID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.ErrNotFound
		}
		return nil, err
	}

	if req.Name != nil {
		client.Name = *req.Name
	}
	if req.Email != nil {
		client.Email = *req.Email
	}
	if req.Phone != nil {
		client.Phone = *req.Phone
	}
	if req.Address != nil {
		client.Address = *req.Address
	}
	if req.Longitude != nil {
		client.Longitude = *req.Longitude
	}
	if req.Latitude != nil {
		client.Latitude = *req.Latitude
	}
	if req.Status != nil {
		client.Status = *req.Status
	}

	if err := s.clientRepo.Update(ctx, client); err != nil {
		return nil, err
	}

	return &dto.ClientResponse{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		Address:   client.Address,
		Longitude: client.Longitude,
		Latitude:  client.Latitude,
		Status:    client.Status,
	}, nil
}

func (s *clientService) DeleteClient(ctx context.Context, clientID int) error {
	client, err := s.clientRepo.FindByID(ctx, clientID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return appErr.ErrNotFound
		}
		return err
	}

	if err := s.clientRepo.Delete(ctx, client); err != nil {
		return err
	}

	return nil
}

func (s *clientService) GetClientByID(ctx context.Context, clientID int) (*dto.ClientResponse, error) {
	client, err := s.clientRepo.FindByID(ctx, clientID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErr.ErrNotFound
		}
		return nil, err
	}

	return &dto.ClientResponse{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		Address:   client.Address,
		Longitude: client.Longitude,
		Latitude:  client.Latitude,
		Status:    client.Status,
	}, nil
}

func (s *clientService) GetClients(ctx context.Context, req dto.GetClientsRequest) (*dto.GetClientsResponse, error) {
	clients, total, err := s.clientRepo.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}

	result := make([]dto.ClientResponse, len(clients))
	for i, c := range clients {
		result[i] = dto.ClientResponse{
			ID:        c.ID,
			Name:      c.Name,
			Email:     c.Email,
			Phone:     c.Phone,
			Address:   c.Address,
			Longitude: c.Longitude,
			Latitude:  c.Latitude,
			Status:    c.Status,
		}
	}

	return &dto.GetClientsResponse{
		Data:   result,
		Limit:  req.Limit,
		Offset: req.Offset,
		Total:  total,
	}, nil
}
